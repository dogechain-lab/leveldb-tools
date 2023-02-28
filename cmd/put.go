package cmd

import (
	"fmt"
	"strings"

	"github.com/dogechain-lab/leveldb-tools/internal/parser"
	"github.com/urfave/cli/v2"
)

var putCmd = &cli.Command{
	Name:   "put",
	Usage:  "put key:value or [key:value]s",
	Action: put,
}

func put(ctx *cli.Context) error {
	db, err := openDB(ctx, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	kvs := ctx.Args()
	if kvs.Len() == 0 {
		return fmt.Errorf("key value pair not set")
	}

	kvargs := kvs.Slice()
	for _, kv := range kvargs {
		kvstr := strings.Split(kv, ":")
		if len(kvstr) < 2 {
			continue
		}

		k, v := kvstr[0], kvstr[1]

		key, err := parser.String2Hex(k)
		if err != nil {
			return fmt.Errorf("key(%s) invalid: %w", k, err)
		}

		value, err := parser.String2Hex(v)
		if err != nil {
			return fmt.Errorf("value(%s) invalid: %w", v, err)
		}

		fmt.Printf("put %#x: %#x to database\n", key, value)

		if err := db.Put(key, value, nil); err != nil {
			return fmt.Errorf("put key(%s) value(%s) failed: %w", k, v, err)
		}
	}

	return nil
}
