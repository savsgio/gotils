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
	return strconv.S2B(V4())
}

// V6 returns a uuid v6 string.
func V6() string {
	return uuid.Must(uuid.NewV6()).String()
}

// V6Bytes returns a uuid v6 byte slice.
func V6Bytes() []byte {
	return strconv.S2B(V6())
}

// V7 returns a uuid v7 string.
func V7() string {
	return uuid.Must(uuid.NewV7()).String()
}

// V7Bytes returns a uuid v7 byte slice.
func V7Bytes() []byte {
	return strconv.S2B(V7())
}
