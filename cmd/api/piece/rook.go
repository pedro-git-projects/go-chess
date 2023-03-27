package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Pawn type represents pawn pieces in a game of chess
type Rook struct {
	color      Color
	moved      bool
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewPawn returns a new pawn instance
func NewRook(color Color, position utils.Coordinate) *Rook {
	return &Rook{
		color:    color,
		moved:    false,
		position: position,
	}
}

// Color is an accessor for the color field
func (r Rook) Color() Color {
	return r.color
}

// Position returns the current piece position
func (r Rook) Position() utils.Coordinate {
	return r.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (r Rook) LegalMoves() []utils.Coordinate {
	return r.legalMoves
}

// Moved sets the pawn moved field to true
func (r *Rook) Moved() {
	r.moved = true
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (r *Rook) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)

	nextRight, ok := board.NRight(r.position, 1)
	for ok && (!board.IsOccupied(nextRight) || !sameColor(nextRight, r, board)) {
		l = append(l, nextRight)
		nextRight, ok = board.NRight(nextRight, 1)
	}

	nextLeft, ok := board.NLeft(r.position, 1)
	for ok && (!board.IsOccupied(nextLeft) || !sameColor(nextLeft, r, board)) {
		l = append(l, nextLeft)
		nextLeft, ok = board.NLeft(nextLeft, 1)
	}

	nextFoward, ok := board.NFoward(r.position, 1)
	for ok && (!board.IsOccupied(nextFoward) || !sameColor(nextFoward, r, board)) {
		l = append(l, nextFoward)
		nextFoward, ok = board.NFoward(nextFoward, 1)
	}

	nextBackward, ok := board.NBackward(r.position, 1)
	for ok && (!board.IsOccupied(nextBackward) || !sameColor(nextBackward, r, board)) {
		l = append(l, nextBackward)
		nextBackward, ok = board.NBackward(nextBackward, 1)
	}

	r.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (r *Rook) Move(to utils.Coordinate, board board) bool {
	r.CalculateLegalMoves(board)
	if utils.Contains(r.legalMoves, to) {
		r.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (r Rook) String() string {
	return fmt.Sprintf("%s rook", r.color)
}
