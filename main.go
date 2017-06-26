package main

import (
	_ "github.com/go-sql-driver/mysql"
	dba "github.com/rpip/dba/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose    = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	configFile = kingpin.Arg("conf", "Config file.").Required().File()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version(dba.Version)
	kingpin.CommandLine.Help = "Anonymize database records."
	kingpin.Parse()

	dba.MustRun(*configFile)
}
