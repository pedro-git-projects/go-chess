package utils

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[T comparable](s []T, v T) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

// Contains reports whether v is present in s.
func Contains[T comparable](s []T, v T) bool {
	return Index(s, v) >= 0
}
