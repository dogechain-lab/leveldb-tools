package byteiter

import "github.com/syndtr/goleveldb/leveldb/util"

// BytesPrefixRange returns key range that satisfy
// - the given prefix, and
// - the given seek position
func BytesPrefixRange(prefix, start []byte) *util.Range {
	r := util.BytesPrefix(prefix)
	r.Start = append(r.Start, start...)

	return r
}
