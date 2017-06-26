package dba

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	dbaConf := "../test-fixtures/sample.conf.hcl"
	if f, err := os.Open(dbaConf); err != nil {
		panic(err)
	} else {
		if _, err := ParseConfig(f); err != nil {
			t.Errorf("expected no errors, but got %v", err)
		}
	}
}
