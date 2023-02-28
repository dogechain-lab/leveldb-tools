package cmd

import (
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:                   "ldb",
	Usage:                  "CRUD leveldb files",
	EnableBashCompletion:   true,
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		putCmd,
		listCmd,
		showCmd,
		deleteCmd,
	},
	Flags: []cli.Flag{
		directoryFlag,
	},
}
