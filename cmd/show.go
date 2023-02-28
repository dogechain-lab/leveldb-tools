package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var showCmd = &cli.Command{
	Name:  "show",
	Usage: "show key: value or [key: value]s, regex supported",
	Flags: []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		fmt.Println("show task: ", cCtx.Args().First())
		return nil
	},
}
