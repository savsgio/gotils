package ulid

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func Test_New(t *testing.T) {
	value := New()

	if _, err := ulid.Parse(value); err != nil {
		t.Errorf("New() == %s, invalid value", value)
	}
}

func Test_NewBytes(t *testing.T) {
	value := NewBytes()

	if _, err := ulid.Parse(string(value)); err != nil {
		t.Errorf("NewBytes() == %s, invalid value", value)
	}
}
