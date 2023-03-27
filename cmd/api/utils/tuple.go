package utils

// Pair represents an ordered pair
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates and returns a pointer to a new instance of the  Pair type
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		First:  first,
		Second: second,
	}
}
