package strings

// IndexOf returns index position in slice from given string
// If value is -1, the string does not found.
func IndexOf(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}

	return -1
}

// Include returns true or false if given string is in slice.
func Include(slice []string, s string) bool {
	return IndexOf(slice, s) >= 0
}

// UniqueAppend appends a string if not exist in the slice.
func UniqueAppend(slice []string, s string) []string {
	if IndexOf(slice, s) < 0 {
		slice = append(slice, s)
	}

	return slice
}
