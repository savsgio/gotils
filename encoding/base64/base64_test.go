// See https://go-review.googlesource.com/c/go/+/37639/1/src/encoding/base64/base64_test.go#18

package base64

import (
	"encoding/base64"
	"testing"
)

type testpair struct {
	decoded, encoded string
}

var pairs = []testpair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},

	// Wikipedia examples
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
}

func testEqual(t *testing.T, msg string, args ...interface{}) { // nolint:thelper
	if args[len(args)-2] != args[len(args)-1] {
		t.Errorf(msg, args...)
	}
}

func TestAppendEncode(t *testing.T) {
	var dbuf []byte

	for _, p := range pairs {
		b := AppendEncode(base64.StdEncoding, dbuf, []byte(p.decoded))
		count := len(b) - len(dbuf)

		testEqual(t, "AppendEncode(%q) = length %v, want %v", p.decoded, count, len(p.encoded))
		testEqual(t, "AppendEncode(%q) = %q, want %q", p.decoded, string(b[len(dbuf):]), p.encoded)

		dbuf = b
	}
}

func TestAppendDecode(t *testing.T) {
	var dbuf []byte

	for _, p := range pairs {
		b, err := AppendDecode(base64.StdEncoding, dbuf, []byte(p.encoded))
		count := len(b) - len(dbuf)

		testEqual(t, "AppendDecode(%q) = error %v, want %v", p.encoded, err, error(nil))
		testEqual(t, "AppendDecode(%q) = length %v, want %v", p.encoded, count, len(p.decoded))
		testEqual(t, "AppendDecode(%q) = %q, want %q", p.encoded, string(b[len(dbuf):]), p.decoded)

		dbuf = b
	}
}

func BenchmarkAppendEncode(b *testing.B) {
	var dbuf []byte

	data := make([]byte, 8192)
	b.SetBytes(int64(len(data)))

	for i := 0; i < b.N; i++ {
		dbuf = AppendEncode(base64.StdEncoding, dbuf[:0], data)
	}
}

func BenchmarkAppendDecode(b *testing.B) {
	var dbuf []byte

	data := AppendEncode(base64.StdEncoding, nil, make([]byte, 8192))
	b.SetBytes(int64(len(data)))

	for i := 0; i < b.N; i++ {
		dbuf, _ = AppendDecode(base64.StdEncoding, dbuf[:0], data)
	}
}
