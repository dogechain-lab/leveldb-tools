package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

type ByteType int

const (
	ByteKey ByteType = iota
	BytePrefix
	ByteAll
)

const (
	expr = `([a-zA-Z0-9\-]+)\{([0-9]+)\}{1}`
)

// Result represents simple regex result
//
// Current support type:
// "prefix" inlcude only one key
// "prefix{keylen}" include all keys start with prefix and key with keylen
// "prefix*" include all keys start with prefix
type Result struct {
	rawStrs string
	btype   ByteType
	prefix  []byte
	keylen  int
}

func (r *Result) Type() ByteType {
	return r.btype
}

func (r *Result) Key() []byte {
	return r.prefix
}

// KeyLen returns key length without prefix
func (r *Result) KeyLen() int {
	switch r.btype {
	case ByteAll, ByteKey:
		return -1
	}

	return r.keylen
}

func (r *Result) ShouldSkip(k []byte) bool {
	switch r.btype {
	case ByteKey:
		return !bytes.EqualFold(r.prefix, k)
	case BytePrefix:
		if len(k) != len(r.prefix)+r.keylen {
			return true
		}
	case ByteAll:
		if len(k) < len(r.prefix) {
			return true
		}
	}

	return !bytes.EqualFold(r.prefix, k[:len(r.prefix)])
}

func ParseKeys(ctx *cli.Context) (*Result, error) {
	s := ctx.Args().Get(0)
	var (
		res = &Result{
			rawStrs: s,
		}
		str    = s
		ishex  bool
		rawkey = s
	)

	if s == "" || s == "*" {
		res.btype = ByteAll
		return res, nil
	}

	if strings.HasPrefix(s, "0x") {
		str = strings.TrimPrefix(s, "0x")
		rawkey = str
		ishex = true
	}

	if strings.HasSuffix(s, "*") {
		str = strings.TrimSuffix(str, "*")
		rawkey = str
		res.btype = ByteAll
	} else if strings.HasSuffix(str, "}") {
		reg, err := regexp.Compile(expr)
		if err != nil {
			return nil, fmt.Errorf("compile regex(%s) failed: %w", expr, err)
		}

		if !reg.MatchString(str) {
			return nil, fmt.Errorf("regex(%s) match failed", expr)
		}

		substrs := reg.FindStringSubmatch(str)
		// the first one is all match
		rawkey = substrs[1]
		res.keylen, err = strconv.Atoi(substrs[2])
		if err != nil {
			return nil, fmt.Errorf("parse str to int failed: %w", err)
		}
		res.btype = BytePrefix
	}

	if ishex {
		res.prefix, _ = hex.DecodeString(rawkey)
	} else {
		res.prefix = []byte(rawkey)
	}

	return res, nil
}
