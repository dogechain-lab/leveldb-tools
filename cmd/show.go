package cmd

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/urfave/cli/v2"
)

var showCmd = &cli.Command{
	Name:   "show",
	Usage:  "show \"key: value\" or [key: value]s, simple regex supported",
	Action: show,
}

func show(ctx *cli.Context) error {
	db, err := openDB(ctx, &opt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	defer db.Close()

	return iterKeys(ctx, db, func(key []byte) error {
		v, err := db.Get(key, nil)
		if err != nil {
			return fmt.Errorf("get key failed: %w", err)
		}
		if len(v) == 0 {
			return fmt.Errorf("value is nil")
		}

		fmt.Printf("%#x: %#x\n", key, v)
		return nil
	}, func(iter iterator.Iterator) error {
		fmt.Printf("%#x: %#x\n", iter.Key(), iter.Value())
		return nil
	})
}
