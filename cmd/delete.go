package cmd

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/urfave/cli/v2"
)

var deleteCmd = &cli.Command{
	Name:    "delete",
	Aliases: []string{"del"},
	Usage:   "delete key or keys, simple regex supported",
	Action:  deleted,
}

func deleted(ctx *cli.Context) error {
	db, err := openDB(ctx, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return iterKeys(ctx, db, func(key []byte) error {
		fmt.Printf("delete key(%#x)\n", key)
		return db.Delete(key, &opt.WriteOptions{
			NoWriteMerge: true, // do not trigger write merge for only 1 key.
		})
	}, func(iter iterator.Iterator) error {
		fmt.Printf("delete key(%#x)\n", iter.Key())
		return db.Delete(iter.Key(), nil)
	})
}
