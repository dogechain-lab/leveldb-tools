package main

import (
	"encoding/hex"
	"fmt"

	"github.com/dogechain-lab/leveldb-tools/internal/byteiter"
	"github.com/dogechain-lab/leveldb-tools/internal/parser"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list out key or keys, regex supported",
	Flags: []cli.Flag{
		directoryFlag,
	},
	Action: list,
}

func list(ctx *cli.Context) error {
	db, err := openDB(ctx, &opt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	defer db.Close()

	res, err := parser.ParseKeys(ctx)
	if err != nil {
		return err
	}

	if res.Type() == parser.ByteKey {
		exists, err := db.Has(res.Key(), nil)
		if err != nil {
			return fmt.Errorf("get key failed: %w", err)
		}
		if !exists {
			return fmt.Errorf("key not exists")
		}

		fmt.Println(hex.EncodeToString(res.Key()))
		return nil
	}

	var (
		iter   iterator.Iterator
		prefix []byte
		keylen int
	)

	var byterange *util.Range
	switch res.Type() {
	case parser.ByteAll:
		byterange = byteiter.BytesPrefixRange(prefix, nil)
	case parser.BytePrefix:
		keylen = res.KeyLen()
		// from zero
		byterange = byteiter.BytesPrefixRange(prefix, make([]byte, keylen))
	}

	iter = db.NewIterator(
		byterange,
		&opt.ReadOptions{
			DontFillCache: true,
		},
	)

	for iter.Next() {
		k := iter.Key()
		if res.ShouldSkip(k) {
			// fmt.Fprintln(os.Stderr, "skip key:", hex.EncodeToString(k))
			continue
		}

		fmt.Println(hex.EncodeToString(k))
	}

	iter.Release()

	return iter.Error()
}
