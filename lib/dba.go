package dba

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/k0kubun/pp"
)

const (
	// Version is the current DBA version
	Version = "v0.1.0"
)

// Anonymizer is the interface implemented by tables for anonimyzing data
type Anonymizer interface {

	// Anonymize should run the actual anonymization
	// It should return an error when it fails or nil otherwise.
	Anonymize(*Database, templateConfig) error

	// GetChangesets returns the set of changes to apply
	GetChangeSets()
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
	Updates map[string]interface{}
}

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
func (tbl *Table) Anonymize(db *Database, tplConfig templateConfig) error {

	columns, updateVals := tbl.GetChangeSets(tplConfig)
	query := fmt.Sprintf("UPDATE %s SET %s", tbl.Name, columns)
	// TODO: print to console that table field is being updated
	fmt.Println(query)
	res, err := db.Exec(query, updateVals...)
	if err != nil {
		log.Fatal(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s : %s, %d rows affected", db.Name, tbl.Name, count)

	return nil
}

// GetChangeSets returns the set of changes to apply to the table
func (tbl *Table) GetChangeSets(tplConfig templateConfig) (string, []interface{}) {

	columns := []string{}
	var changesets []interface{}

	for k, v := range tbl.Updates {
		result, err := mustEvalTemplate(v, tplConfig)
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
func (db *Database) Run(tplConfig templateConfig) {

	if err := db.Connect(); err != nil {
		db.Close()
		log.Fatalf("%s %v", db.Name, err)
	}

	for _, tbl := range db.Tables {
		if tbl.isAnonymizable() {
			if err := tbl.Anonymize(db, tplConfig); err != nil {
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

// Run triggers the anonimyzers on the database tables
func Run(conf *Config) {
	pp.Print(conf)
	registerBuiltins(conf)

	for _, db := range conf.Databases {
		db.Run(conf.templateConfig)
	}
}
