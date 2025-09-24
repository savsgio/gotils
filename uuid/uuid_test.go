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

func Test_V6(t *testing.T) {
	value := V6()

	if _, err := uuid.Parse(value); err != nil {
		t.Errorf("V6() == %s, invalid value", value)
	}
}

func Test_V6Bytes(t *testing.T) {
	value := V6Bytes()

	if _, err := uuid.ParseBytes(value); err != nil {
		t.Errorf("V6() == %s, invalid value", value)
	}
}

func Test_V7(t *testing.T) {
	value := V7()

	if _, err := uuid.Parse(value); err != nil {
		t.Errorf("V7() == %s, invalid value", value)
	}
}

func Test_V7Bytes(t *testing.T) {
	value := V7Bytes()

	if _, err := uuid.ParseBytes(value); err != nil {
		t.Errorf("V7() == %s, invalid value", value)
	}
}
