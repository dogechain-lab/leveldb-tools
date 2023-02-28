package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var putCmd = &cli.Command{
	Name:  "put",
	Usage: "put key value or [key value]s",
	Flags: []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		fmt.Println("list task: ", cCtx.Args().First(), cCtx.Args().Get(2))
		return nil
	},
}
