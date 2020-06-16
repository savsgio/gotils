package gotils

// StringSliceIndexOf returns index position in slice from given string
// If value is -1, the string does not found.
func StringSliceIndexOf(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}

	return -1
}

// StringSliceInclude returns true or false if given string is in slice.
func StringSliceInclude(slice []string, s string) bool {
	return StringSliceIndexOf(slice, s) >= 0
}

// StringUniqueAppend appends a string if not exist in the slice.
func StringUniqueAppend(slice []string, s string) []string {
	if StringSliceIndexOf(slice, s) < 0 {
		slice = append(slice, s)
	}

	return slice
}
