package gotils

import (
	"github.com/savsgio/gotils/bytes"
	"github.com/savsgio/gotils/strconv"
	"github.com/savsgio/gotils/strings"
)

// Deprecated: Use
//
// import "github.com/savsgio/gotils/bytes"
//
// bytes.Extend(b, needLen).
func ExtendByteSlice(b []byte, needLen int) []byte {
	return bytes.Extend(b, needLen)
}

// Deprecated: Use
//
// import "github.com/savsgio/gotils/bytes"
//
// bytes.Rand(dst).
func RandBytes(dst []byte) []byte {
	return bytes.Rand(dst)
}

// Deprecated: Use
//
// import "github.com/savsgio/gotils/strconv"
//
// strconv.B2S(b).
func B2S(b []byte) string {
	return strconv.B2S(b)
}

// Deprecated: Use
//
// import "github.com/savsgio/gotils/strconv"
//
// strconv.S2B(s).
func S2B(s string) []byte {
	return strconv.S2B(s)
}

// Deprecated: Use
//
// import "github.com/savsgio/gotils/strings"
//
// strings.IndexOf(slice, s).
func StringSliceIndexOf(slice []string, s string) int {
	return strings.IndexOf(slice, s)
}

// Deprecated: Use
//
// import "github.com/savsgio/gotils/strings"
//
// strings.Include(slice, s).
func StringSliceInclude(slice []string, s string) bool {
	return strings.Include(slice, s)
}

// Deprecated: Use
//
// import "github.com/savsgio/gotils/strings"
//
// strings.UniqueAppend(slice, s).
func StringUniqueAppend(slice []string, s string) []string {
	return strings.UniqueAppend(slice, s)
}
