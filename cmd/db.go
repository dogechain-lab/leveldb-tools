package main

import (
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
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
