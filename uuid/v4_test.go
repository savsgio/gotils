package uuid

import (
	"testing"

	"github.com/google/uuid"
)

func Test_V4(t *testing.T) {
	value := V4()

	if _, err := uuid.Parse(value); err != nil {
		t.Errorf("V4() == %s, invalid value", value)
	}
}

func Test_V4Bytes(t *testing.T) {
	value := V4Bytes()

	if _, err := uuid.ParseBytes(value); err != nil {
		t.Errorf("V4() == %s, invalid value", value)
	}
}
