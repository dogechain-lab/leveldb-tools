package main

import (
	"fmt"
	"os"

	"github.com/dogechain-lab/leveldb-tools/internal/byteiter"
	"github.com/dogechain-lab/leveldb-tools/internal/parser"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/urfave/cli/v2"
)

func openDB(ctx *cli.Context, opts *opt.Options) (*leveldb.DB, error) {
	dbpath := expandPath(ctx.Path(directoryFlag.Name))
	if exist, err := fileExists(dbpath); err != nil || !exist {
		return nil, fmt.Errorf("check file(%s) failed: %w", dbpath, err)
	}

	db, err := leveldb.OpenFile(dbpath, opts)
	if err != nil {
		return nil, fmt.Errorf("open DB(%s) failed: %w", dbpath, err)
	}

	return db, nil
}

func fileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}

	return true, nil
}

func iterKeys(ctx *cli.Context,
	db *leveldb.DB,
	keyFn func(key []byte) error,
	iterFn func(iter iterator.Iterator) error,
) error {
	res, err := parser.ParseKeys(ctx)
	if err != nil {
		return err
	}

	if res.Type() == parser.ByteKey {
		if keyFn != nil {
			return keyFn(res.Key())
		}
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
	defer iter.Release()

	for iter.Next() {
		k := iter.Key()
		if res.ShouldSkip(k) {
			// fmt.Fprintln(os.Stderr, "skip key:", hex.EncodeToString(k))
			continue
		}

		if iterFn != nil {
			if err := iterFn(iter); err != nil {
				return err
			}
		}
	}

	return iter.Error()
}
