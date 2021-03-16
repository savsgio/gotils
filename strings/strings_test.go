package strings

import (
	"reflect"
	"testing"
)

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

	result = UniqueAppend(slice2, slice...)

	if len(result) == len(slice) {
		t.Errorf("StringUniqueAppend() == %v, want %v", result, slice2)
	}
}

func Test_Reverse(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}
	reversedSlice := []string{"cache", "http", "fast", "kratgo"}

	result := Reverse(slice)

	if !reflect.DeepEqual(result, reversedSlice) {
		t.Errorf("Reverse() == %v, want %v", result, reversedSlice)
	}
}

func Test_Copy(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	result := Copy(slice)

	if !reflect.DeepEqual(result, slice) {
		t.Errorf("Copy() == %v, want %v", result, slice)
	}
}

func Test_Equal(t *testing.T) {
	slice1 := []string{"kratgo", "fast", "http", "cache"}
	slice2 := []string{"kratgo", "fast", "http", "cache"}

	if !Equal(slice1, slice2) {
		t.Errorf("Equal() == %v, want %v", false, true)
	}

	slice1 = []string{"kratgo", "fast", "http", "cache"}
	slice2 = []string{"kratgo", "fast", "http"}

	if Equal(slice1, slice2) {
		t.Errorf("Equal() == %v, want %v", true, false)
	}

	slice1 = []string{"kratgo", "fast", "http", "cache"}
	slice2 = []string{"kratgo", "fast", "rpc", "cache"}

	if Equal(slice1, slice2) {
		t.Errorf("Equal() == %v, want %v", true, false)
	}
}
