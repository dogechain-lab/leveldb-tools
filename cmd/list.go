package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list out key or keys, regex supported",
	Flags:   []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		fmt.Println("list task: ", cCtx.Args().First())
		return nil
	},
}
