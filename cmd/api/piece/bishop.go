package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Bishop type represents the bishop piece in a game of chess
type Bishop struct {
	color      Color
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewBishop returns a new king instance
func NewBishop(color Color, position utils.Coordinate) *Bishop {
	return &Bishop{
		color:    color,
		position: position,
	}
}

// Color is an accessor for the color field
func (b Bishop) Color() Color {
	return b.color
}

// Position returns the current piece position
func (b Bishop) Position() utils.Coordinate {
	return b.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (b Bishop) LegalMoves() []utils.Coordinate {
	return b.legalMoves
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
// TODO check if diagonals are being calculate correctly
func (b *Bishop) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)

	nextfrd, ok := board.NthFowardRightDiagonal(b.position, 1)
	for ok && (!board.IsOccupied(nextfrd) || !sameColor(nextfrd, b, board)) {
		l = append(l, nextfrd)
		nextfrd, ok = board.NthFowardRightDiagonal(nextfrd, 1)
	}

	nextfld, ok := board.NthFowardLeftDiagonal(b.position, 1)
	for ok && (!board.IsOccupied(nextfld) || !sameColor(nextfld, b, board)) {
		l = append(l, nextfld)
		nextfld, ok = board.NthFowardLeftDiagonal(nextfld, 1)
	}

	nextbrd, ok := board.NthBackwardRightDiagonal(b.position, 1)
	for ok && (!board.IsOccupied(nextbrd) || !sameColor(nextbrd, b, board)) {
		l = append(l, nextbrd)
		nextbrd, ok = board.NthBackwardRightDiagonal(nextbrd, 1)
	}

	nextbld, ok := board.NthBackwardLeftDiagonal(b.position, 1)
	for ok && (!board.IsOccupied(nextbld) || !sameColor(nextbld, b, board)) {
		l = append(l, nextbld)
		nextbld, ok = board.NthBackwardLeftDiagonal(nextbld, 1)
	}

	b.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (b *Bishop) Move(to utils.Coordinate, board board) bool {
	b.CalculateLegalMoves(board)
	if utils.Contains(b.legalMoves, to) {
		b.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (b Bishop) String() string {
	return fmt.Sprintf("%s bishop", b.color)
}
