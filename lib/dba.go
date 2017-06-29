package dba

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	// Version is the current DBA version
	Version = "v0.1.0"
)

// Anonymizer is the interface implemented by tables for anonimyzing data
type Anonymizer interface {

	// Anonymize should run the actual anonymization
	// It should return an error when it fails or nil otherwise.
	Anonymize(*Database, EvalConfig) error

	// GetChangesets returns the set of changes to apply
	GetChangeSets(EvalConfig)
}

type meta struct {
	Name     string
	Scope    string
	Metadata map[string]interface{}
}

// GetMeta retrieves value of key. Panics if key is missing
func (m *meta) MustGetMeta(key string) interface{} {
	if val, ok := m.Metadata[key]; ok {
		return val
	}
	log.Fatalf("Expected config '%s' in '%s' %s", key, m.Name, m.Scope)
	return nil
}

func (m *meta) SetMeta(key string, val interface{}) {
	m.Metadata[key] = val
}

// Table represents the table node in the config file
type Table struct {
	meta
	Updates       map[string]interface{}
	updateCounter int
}

// TableRow represents a row in a table
type TableRow map[string]interface{}

// NewTable creates a new Table struct
func NewTable(tblName string) *Table {
	return &Table{
		meta: meta{
			Name:     tblName,
			Scope:    "table",
			Metadata: make(map[string]interface{}),
		},
		Updates: make(map[string]interface{}),
	}
}

func (tbl *Table) isAnonymizable() bool {
	return len(tbl.Updates) > 0
}

// Anonymize does the actual table anonymization
func (tbl *Table) Anonymize(db *Database, Ctx EvalConfig) error {
	rows, err := db.Query(fmt.Sprintf("SELECT * from %s", tbl.Name))
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		row := tbl.mapScan(rows, cols)
		// TODO: handle error
		tbl.runUpdate(db, row, Ctx)
	}
	log.Printf("%s : %s, %d rows affected", db.Name, tbl.Name, tbl.updateCounter)
	return nil
}

func (tbl *Table) mapScan(rows *sql.Rows, cols []string) TableRow {
	count := len(cols)
	columns := make([]interface{}, count)
	values := make([]interface{}, count)

	for i := range cols {
		values[i] = &columns[i]
	}
	if err := rows.Scan(values...); err != nil {
		log.Fatal(err)
	}

	row := make(TableRow)

	for i, col := range cols {
		var v interface{}
		val := columns[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		row[col] = v
	}

	return row
}

// runUpdate runs the actual row update query
func (tbl *Table) runUpdate(db *Database, row TableRow, Ctx EvalConfig) error {
	fields, updateVals := tbl.GetChangeSets(Ctx)
	query := fmt.Sprintf("UPDATE %s SET %s", tbl.Name, fields)

	var primaryKey string
	if val, ok := tbl.Metadata["primary_key"]; ok {
		primaryKey = val.(string)
	} else {
		primaryKey = defaultPrimaryKey
	}

	query += fmt.Sprintf(" WHERE %s = ?", primaryKey)
	updateVals = append(updateVals, row[primaryKey])
	_, err := db.Exec(query, updateVals...)
	tbl.updateCounter++
	return err
}

// GetChangeSets returns the set of changes to apply to the table
func (tbl *Table) GetChangeSets(Ctx EvalConfig) (string, []interface{}) {

	columns := []string{}
	var changesets []interface{}

	for k, v := range tbl.Updates {
		result, err := EvalTemplate(v, Ctx)
		if err != nil {
			log.Fatal(TemplateError{
				err:     err,
				tblName: tbl.Name,
				field:   k,
				input:   v,
			})
		}
		columns = append(columns, k)
		changesets = append(changesets, result)
	}
	// strings.Join does not add separator to the last element
	columnsStr := strings.Join(columns[:], " = ?, ") + " = ?"
	return columnsStr, changesets
}

// Database represents the 'db' node in the config file
type Database struct {
	meta
	Tables []*Table
	*sql.DB
}

// NewDB creates a new Database struct
func NewDB(dbName string) *Database {
	return &Database{
		meta: meta{
			Name:     dbName,
			Scope:    "database",
			Metadata: make(map[string]interface{}),
		},
	}
}

// Connect creates the DB connection
func (db *Database) Connect() error {
	dsn := db.MustGetMeta("dsn").(string)
	driver := db.MustGetMeta("type").(string)
	db.DB, _ = sql.Open(driver, dsn)

	err := db.Ping()
	return err
}

// Run establishes DB connection and anonymizes the tables
func (db *Database) Run(Ctx EvalConfig) {

	if err := db.Connect(); err != nil {
		db.Close()
		log.Fatalf("%s %v", db.Name, err)
	}

	for _, tbl := range db.Tables {
		if tbl.isAnonymizable() {
			if err := tbl.Anonymize(db, Ctx); err != nil {
				db.Close()
				log.Fatal(err)
			}
		}
	}
}

// Close ensures DB connection is closed
func (db *Database) Close() {
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
}

// MustRun triggers the anonimyzers on the database tables. If an error occurs, execution stops.
func MustRun(config io.ReadWriter) {

	conf, err := ParseConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	// pp.Print(conf)
	conf.EvalConfig = buildTemplateContext()

	for _, db := range conf.Databases {
		db.Run(conf.EvalConfig)
	}
}
