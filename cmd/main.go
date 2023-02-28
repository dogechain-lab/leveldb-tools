package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:  "ldb",
	Usage: "read/write leveldb files",
	Commands: []*cli.Command{
		putCmd,
		listCmd,
		showCmd,
		deleteCmd,
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
