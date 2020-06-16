package gotils

import "testing"

func Test_StringSliceIndexOf(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	s := "fast"
	if i := StringSliceIndexOf(slice, s); i < 0 {
		t.Errorf("stringSliceIndexOf() = %v, want %v", i, 2)
	}

	s = "slow"
	if i := StringSliceIndexOf(slice, s); i > -1 {
		t.Errorf("stringSliceIndexOf() = %v, want %v", i, -1)
	}
}

func Test_StringSliceInclude(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	s := "fast"
	if ok := StringSliceInclude(slice, s); !ok {
		t.Errorf("stringSliceInclude() = %v, want %v", ok, true)
	}

	s = "slow"
	if ok := StringSliceInclude(slice, s); ok {
		t.Errorf("stringSliceInclude() = %v, want %v", ok, false)
	}
}

func Test_StringSliceInc(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	result := StringUniqueAppend(slice, slice[0])

	if len(result) != len(slice) {
		t.Errorf("StringUniqueAppend() == %v, want %v", result, slice)
	}

	s := "unique"
	slice2 := append(slice, s)

	result = StringUniqueAppend(slice2, slice[0])

	if len(result) == len(slice) {
		t.Errorf("StringUniqueAppend() == %v, want %v", result, slice2)
	}
}
