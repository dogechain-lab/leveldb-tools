package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var deleteCmd = &cli.Command{
	Name:    "delete",
	Aliases: []string{"del"},
	Usage:   "delete key or keys, regex supported",
	Flags:   []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		fmt.Println("delete task: ", cCtx.Args().First())
		return nil
	},
}
