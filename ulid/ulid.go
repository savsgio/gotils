package ulid

import (
	"github.com/oklog/ulid/v2"
	"github.com/savsgio/gotils/strconv"
)

// New returns a new ULID string.
func New() string {
	return ulid.Make().String()
}

// NewBytes returns a new ULID byte slice.
func NewBytes() []byte {
	return strconv.S2B(New())
}
