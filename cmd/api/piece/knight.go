package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Knight type represents the king piece in a game of chess
type Knight struct {
	color      Color
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewKnight returns a new king instance
func NewKnight(color Color, position utils.Coordinate) *Knight {
	return &Knight{
		color:    color,
		position: position,
	}
}

// Color is an accessor for the color field
func (k Knight) Color() Color {
	return k.color
}

// Position returns the current piece position
func (k Knight) Position() utils.Coordinate {
	return k.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (k Knight) LegalMoves() []utils.Coordinate {
	return k.legalMoves
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (k *Knight) CalculateLegalMoves(board board) {
	l := []utils.Coordinate{}
	temp := board.CalcAllLs(k.position)
	for _, v := range temp {
		if !sameColor(v, k, board) {
			l = append(l, v)
		}
	}
	k.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (k *Knight) Move(to utils.Coordinate, board board) bool {
	k.CalculateLegalMoves(board)
	if utils.Contains(k.legalMoves, to) {
		k.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (k Knight) String() string {
	return fmt.Sprintf("%s knight", k.color)
}
