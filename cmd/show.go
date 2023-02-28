package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/urfave/cli/v2"
)

var showCmd = &cli.Command{
	Name:  "show",
	Usage: "show key: value or [key: value]s, regex supported",
	Flags: []cli.Flag{
		directoryFlag,
	},
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

		fmt.Printf("%s: %s\n", hex.EncodeToString(key), hex.EncodeToString(v))
		return nil
	}, func(iter iterator.Iterator) error {
		fmt.Printf("%s: %s\n", hex.EncodeToString(iter.Key()), hex.EncodeToString(iter.Value()))
		return nil
	})
}
