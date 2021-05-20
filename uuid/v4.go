package uuid

import (
	"github.com/google/uuid"
	"github.com/savsgio/gotils/strconv"
)

// V4 returns a uuid v4 string.
func V4() string {
	return uuid.New().String()
}

// V4Bytes returns a uuid v4 byte slice.
func V4Bytes() []byte {
	return strconv.S2B(uuid.New().String())
}
