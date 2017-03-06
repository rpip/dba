package dba

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

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
	Anonymize(*Database) error

	// GetChangesets returns the set of changes to apply
	GetChangeSets()
}

type meta struct {
	Name     string
	Scope    string
	Metadata map[string]interface{}
}

// Database represents the db node in the config file
type Database struct {
	meta
	Tables []*Table
	*sql.DB
}

// Table represents the table node in the config file
type Table struct {
	meta
	Updates map[string]interface{}
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

// Connect creates the DB connection
func (db *Database) Connect() error {
	dsn := db.MustGetMeta("dsn").(string)
	driver := db.MustGetMeta("type").(string)
	db.DB, _ = sql.Open(driver, dsn)

	err := db.Ping()
	return err
}

// Anonymize does the actual table anonymization
func (tbl *Table) Anonymize(db *Database) error {

	columns, updateVals := tbl.GetChangeSets()
	query := fmt.Sprintf("UPDATE %s SET %s", tbl.Name, columns)
	fmt.Print(query)
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
func (tbl *Table) GetChangeSets() (string, []interface{}) {
	columns := []string{}
	changesets := make([]interface{}, len(tbl.Updates))
	for k, v := range tbl.Updates {
		columns = append(columns, k)
		if v, err := mustEvalTemplate(v); err != nil {
			log.Fatal(TemplateError{
				err:     err,
				tblName: tbl.Name,
				field:   k,
				input:   v,
			})
		}
		changesets = append(changesets, v)
	}
	// strings.Join does not add separator to the last element
	columnsStr := strings.Join(columns[:], " = ?, ") + " = ?"
	return columnsStr, changesets
}

func mustEvalTemplate(v interface{}) (interface{}, error) {
	switch v := v.(type) {
	case string:
		return evalTemplate(v)
	default:
		return v, nil
	}
}

func (tbl *Table) isAnonymizable() bool {
	return len(tbl.Updates) > 0
}

// RunAnonymizers triggers the anonimyzers on the database tables
func RunAnonymizers(conf Config) {
	pp.Print(conf)
	for _, db := range conf.Databases {
		func(db *Database) {
			var wg sync.WaitGroup
			wg.Add(len(db.Tables))
			defer func() {
				if err := db.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			if err := db.Connect(); err != nil {
				log.Fatalf("%s %v", db.Name, err)
			}

			for _, tbl := range db.Tables {
				if tbl.isAnonymizable() {
					// TODO: fatal error: concurrent map read and map write
					go func(tbl *Table, wg *sync.WaitGroup) {
						if err := tbl.Anonymize(db); err != nil {
							log.Fatal(err)
						}
						wg.Done()
					}(tbl, &wg)
				}
			}
			wg.Wait()
		}(db)
	}
}
