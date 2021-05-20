package bytes

import (
	"math/rand"
	"time"

	"github.com/savsgio/gotils/strconv"
)

var src = rand.NewSource(time.Now().UnixNano())

// Rand returns dst with a string random bytes
// Make sure that dst has the length you need.
func Rand(dst []byte) []byte {
	n := len(dst)

	for i, cache, remain := n-1, src.Int63(), charsetIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), charsetIdxMax
		}

		if idx := int(cache & charsetIdxMask); idx < len(charset) {
			dst[i] = charset[idx]
			i--
		}

		cache >>= charsetIdxBits
		remain--
	}

	return dst
}

// Copy returns a copy of byte slice in a new pointer.
func Copy(b []byte) []byte {
	return []byte(strconv.B2S(b))
}

// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
	return strconv.B2S(a) == strconv.B2S(b)
}

// Extend extends b to needLen bytes.
func Extend(b []byte, needLen int) []byte {
	b = b[:cap(b)]
	if n := needLen - cap(b); n > 0 {
		b = append(b, make([]byte, n)...)
	}

	return b[:needLen]
}

// Prepend prepends bytes into a given byte slice.
func Prepend(dst []byte, src ...byte) []byte {
	dstLen := len(dst)
	srcLen := len(src)

	dst = Extend(dst, dstLen+srcLen)
	copy(dst[srcLen:], dst[:dstLen])
	copy(dst[:srcLen], src)

	return dst
}

// PrependString prepends a string into a given byte slice.
func PrependString(dst []byte, src string) []byte {
	return Prepend(dst, strconv.S2B(src)...)
}
