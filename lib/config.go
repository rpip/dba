package dba

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hil"
)

// defaultPrimaryKey Table primary key
const defaultPrimaryKey = "id"

// hclConfig is map structure to hold parsed config
type hclConfig map[string][]map[string][]map[string]interface{}

type templateConfig struct {
	*hil.EvalConfig
}

// Config is the HCL config parsed into databases and tables
type Config struct {
	templateConfig
	Databases []*Database
}

// ConfigParseError denotes failing to parse configuration file.
type ConfigParseError struct {
	err error
}

// Error returns the formatted configuration error.
func (pe ConfigParseError) Error() string {
	return fmt.Sprintf("While parsing config: %s", pe.err.Error())
}

func loadConfig(configFile io.ReadWriter) (hclConfig, error) {
	var obj hclConfig

	hclText, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, ConfigParseError{err}
	}

	err = hcl.Decode(&obj, string(hclText))
	if err != nil {
		return nil, ConfigParseError{err}
	}

	return obj, nil
}

// ParseConfig parses the DBA config
func ParseConfig(configFile io.ReadWriter) (*Config, error) {

	dbaConfig := &Config{}
	conf, err := loadConfig(configFile)

	if err != nil {
		return nil, err
	}

	for _, dbNode := range conf["db"] {
		for dbName, v := range dbNode {
			db := NewDB(dbName)

			for k, dbNodeVal := range v[0] {
				if k == "table" {
					// build table
					buildTable(db, dbNodeVal)
				} else {
					db.SetMeta(k, dbNodeVal)
				}
			}
			dbaConfig.Databases = append(dbaConfig.Databases, db)
		}
	}
	return dbaConfig, nil
}

func buildTable(db *Database, dbNodeVal interface{}) {
	dbNodeVals := dbNodeVal.([]map[string]interface{})
	for _, tblNodes := range dbNodeVals {
		for tblName, tblNode := range tblNodes {
			tbl := NewTable(tblName)
			if tblNodeVals, ok := tblNode.([]map[string]interface{}); ok {
				for _, tblNodeVal := range tblNodeVals {
					for tblK, tblV := range tblNodeVal {
						if tblK == "updates" {
							Updates := tblV.([]map[string]interface{})[0]
							tbl.Updates = Updates
						} else {
							tbl.SetMeta(tblK, tblV)
						}
					}
					db.Tables = append(db.Tables, tbl)
				}
			}
		}
	}
}
