package base64

import (
	"encoding/base64"

	"github.com/savsgio/gotils/bytes"
)

// AppendEncode appends encoded src to dst using the encoding enc
// and returns the extended dst.
//
// The encoding pads the output to a multiple of 4 bytes,
// so AppendEncode is not appropriate for use on individual blocks
// of a large data stream. Use NewEncoder() instead.
//
// See https://go-review.googlesource.com/c/go/+/37639/1/src/encoding/base64/base64.go#149
func AppendEncode(enc *base64.Encoding, dst, src []byte) []byte {
	sliceLen := enc.EncodedLen(len(src)) + len(dst)
	b := bytes.Extend(dst, sliceLen)
	enc.Encode(b[len(dst):], src)

	return b
}

// AppendDecode appends decoded src to dst using the encoding enc
// and returns the extended dst
//
// If src contains invalid base64 data, it will return the CorruptInputError.
// New line characters (\r and \n) are ignored.
//
// See https://go-review.googlesource.com/c/go/+/37639/1/src/encoding/base64/base64.go#372
func AppendDecode(enc *base64.Encoding, dst, src []byte) ([]byte, error) {
	sliceLen := enc.DecodedLen(len(src)) + len(dst)
	b := bytes.Extend(dst, sliceLen)
	n, err := enc.Decode(b[len(dst):], src)

	return b[:len(dst)+n], err // nolint:wrapcheck
}
