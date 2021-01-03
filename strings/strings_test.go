package strings

import "testing"

func Test_IndexOf(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	s := "fast"
	if i := IndexOf(slice, s); i < 0 {
		t.Errorf("stringSliceIndexOf() = %v, want %v", i, 2)
	}

	s = "slow"
	if i := IndexOf(slice, s); i > -1 {
		t.Errorf("stringSliceIndexOf() = %v, want %v", i, -1)
	}
}

func Test_Include(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	s := "fast"
	if ok := Include(slice, s); !ok {
		t.Errorf("stringSliceInclude() = %v, want %v", ok, true)
	}

	s = "slow"
	if ok := Include(slice, s); ok {
		t.Errorf("stringSliceInclude() = %v, want %v", ok, false)
	}
}

func Test_UniqueAppend(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	result := UniqueAppend(slice, slice[0])

	if len(result) != len(slice) {
		t.Errorf("StringUniqueAppend() == %v, want %v", result, slice)
	}

	s := "unique"
	slice2 := append(slice, s)

	result = UniqueAppend(slice2, slice[0])

	if len(result) == len(slice) {
		t.Errorf("StringUniqueAppend() == %v, want %v", result, slice2)
	}
}
