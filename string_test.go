package gotils

import "testing"

func Test_StringSliceIndexOf(t *testing.T) {
	array := []string{"kratgo", "fast", "http", "cache"}

	s := "fast"
	if i := StringSliceIndexOf(array, s); i < 0 {
		t.Errorf("stringSliceIndexOf() = %v, want %v", i, 2)
	}

	s = "slow"
	if i := StringSliceIndexOf(array, s); i > -1 {
		t.Errorf("stringSliceIndexOf() = %v, want %v", i, -1)
	}
}

func Test_StringSliceInclude(t *testing.T) {
	array := []string{"kratgo", "fast", "http", "cache"}

	s := "fast"
	if ok := StringSliceInclude(array, s); !ok {
		t.Errorf("stringSliceInclude() = %v, want %v", ok, true)
	}

	s = "slow"
	if ok := StringSliceInclude(array, s); ok {
		t.Errorf("stringSliceInclude() = %v, want %v", ok, false)
	}
}
