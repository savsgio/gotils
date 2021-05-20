package strings

import (
	"reflect"
	"testing"
)

func Test_Copy(t *testing.T) {
	str := "cache"
	result := Copy(str)

	if reflect.ValueOf(&str).Pointer() == reflect.ValueOf(&result).Pointer() {
		t.Errorf("Copy() returns the same pointer (source == %p | result == %p)", &str, &result)
	}
}

func Test_CopySlice(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}

	result := CopySlice(slice)

	if !reflect.DeepEqual(result, slice) {
		t.Errorf("CopySlice() == %v, want %v", result, slice)
	}
}

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
	slice2 := CopySlice(slice)

	result := UniqueAppend(slice, slice[0])
	if len(result) != len(slice) {
		t.Errorf("UniqueAppend() == %v, want %v", result, slice)
	}

	slice2 = append(slice2, "unique")

	result = UniqueAppend(slice, slice2...)
	if len(result) == len(slice) {
		t.Errorf("UniqueAppend() == %v, want %v", result, slice2)
	}
}

func Test_EqualSlices(t *testing.T) {
	slice1 := []string{"kratgo", "fast", "http", "cache"}
	slice2 := []string{"kratgo", "fast", "http", "cache"}

	if !EqualSlices(slice1, slice2) {
		t.Errorf("EqualSlices() == %v, want %v", false, true)
	}

	slice1 = []string{"kratgo", "fast", "http", "cache"}
	slice2 = []string{"kratgo", "fast", "http"}

	if EqualSlices(slice1, slice2) {
		t.Errorf("EqualSlices() == %v, want %v", true, false)
	}

	slice1 = []string{"kratgo", "fast", "http", "cache"}
	slice2 = []string{"kratgo", "fast", "rpc", "cache"}

	if EqualSlices(slice1, slice2) {
		t.Errorf("EqualSlices() == %v, want %v", true, false)
	}
}

func Test_ReverseSlice(t *testing.T) {
	slice := []string{"kratgo", "fast", "http", "cache"}
	reversedSlice := []string{"cache", "http", "fast", "kratgo"}

	result := ReverseSlice(slice)

	if !EqualSlices(result, reversedSlice) {
		t.Errorf("ReverseSlice() == %v, want %v", result, reversedSlice)
	}
}
