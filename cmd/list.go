package cmd

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list out key or keys, simple regex supported",
	Action:  list,
}

func list(ctx *cli.Context) error {
	db, err := openDB(ctx, &opt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	defer db.Close()

	return iterKeys(ctx, db, func(key []byte) error {
		exists, err := db.Has(key, nil)
		if err != nil {
			return fmt.Errorf("get key failed: %w", err)
		}
		if !exists {
			return fmt.Errorf("key not exists")
		}

		fmt.Printf("%#x\n", key)
		return nil
	}, func(iter iterator.Iterator) error {
		fmt.Printf("%#x\n", iter.Key())
		return nil
	})
}
