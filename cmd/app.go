package cmd

import (
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:                   "ldb",
	Usage:                  "CRUD leveldb files",
	EnableBashCompletion:   true,
	UseShortOptionHandling: true,
	HideVersion:            true, // we create a command for this
	Commands: []*cli.Command{
		putCmd,
		listCmd,
		showCmd,
		deleteCmd,
		versionCmd,
	},
	Flags: []cli.Flag{
		directoryFlag,
	},
}
