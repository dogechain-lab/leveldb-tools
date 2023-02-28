package main

import (
	"log"
	"os"

	"github.com/dogechain-lab/leveldb-tools/cmd"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
