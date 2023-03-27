package utils

import (
	"fmt"
	"strconv"
	"unicode"
)

// Coordinate represents an (y,x) chess board coordinate
type Coordinate Pair[rune, int]

// NewCoordinate returns a new instance of the
// Coordinate type
func NewCoordinate(y rune, x int) Coordinate {
	return Coordinate{y, x}
}

// CoordFromStr takes a string of the form a1
// and retuns a coordinate if it is possible
// and an error otherwise
func CoordFromStr(s string) (Coordinate, error) {
	coord := &Coordinate{}
	for pos, char := range s {
		if unicode.IsDigit(char) {
			str := s[:pos]
			char := []rune(str)
			num, err := strconv.Atoi(s[pos:])
			if err != nil {
				return Coordinate{}, err
			}
			coord.First = char[0]
			coord.Second = num
		}
	}
	return *coord, nil
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%c, %d)", c.First, c.Second)
}
