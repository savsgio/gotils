package strings

import (
	"slices"

	"github.com/savsgio/gotils/strconv"
)

// Copy returns a copy of string in a new pointer.
func Copy(s string) string {
	return string(strconv.S2B(s))
}

// UniqueAppend appends a string if not exist in the slice.
func UniqueAppend(slice []string, s ...string) []string {
	for i := range s {
		if IndexOf(slice, s[i]) != -1 {
			continue
		}

		slice = append(slice, s[i])
	}

	return slice
}

// CopySlice returns a copy of the slice.
//
// Deprecated: Use slices.Clone instead, it will be removed in the future.
func CopySlice(slice []string) []string {
	return slices.Clone(slice)
}

// IndexOf returns index position in slice from given string
// If value is -1, the string does not found.
//
// Deprecated: Use slices.Index instead, it will be removed in the future.
func IndexOf(s []string, v string) int {
	return slices.Index(s, v)
}

// Include returns true or false if given string is in slice.
//
// Deprecated: Use slices.Contains instead, it will be removed in the future.
func Include(s []string, v string) bool {
	return slices.Contains(s, v)
}

// EqualSlices checks if the slices are equal.
//
// Deprecated: Use slices.Equal instead, it will be removed in the future.
func EqualSlices(s1, s2 []string) bool {
	return slices.Equal(s1, s2)
}

// ReverseSlice reverses a string slice.
//
// Deprecated: Use slices.Reverse instead, it will be removed in the future.
func ReverseSlice(s []string) []string {
	slices.Reverse(s)

	return s
}
