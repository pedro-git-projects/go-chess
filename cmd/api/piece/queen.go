package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Queen type represents the queen piece in a game of chess
type Queen struct {
	color      Color
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewQueen returns a new queen instance
func NewQueen(color Color, position utils.Coordinate) *Queen {
	return &Queen{
		color:    color,
		position: position,
	}
}

// Color is an accessor for the color field
func (q Queen) Color() Color {
	return q.color
}

// Position returns the current piece position
func (q Queen) Position() utils.Coordinate {
	return q.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (q Queen) LegalMoves() []utils.Coordinate {
	return q.legalMoves
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (q *Queen) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)

	foward, ok := board.NFoward(q.position, 1)
	for ok && (!board.IsOccupied(foward) || !sameColor(foward, q, board)) {
		l = append(l, foward)
		foward, ok = board.NFoward(foward, 1)
	}

	backward, ok := board.NBackward(q.position, 1)
	for ok && (!board.IsOccupied(backward) || !sameColor(backward, q, board)) {
		l = append(l, backward)
		backward, ok = board.NBackward(backward, 1)
	}

	// lateral left
	left, ok := board.NLeft(q.position, 1)
	for ok && (!board.IsOccupied(left) || !sameColor(left, q, board)) {
		l = append(l, left)
		left, ok = board.NLeft(left, 1)
	}

	right, ok := board.NRight(q.position, 1)
	for ok && (!board.IsOccupied(right) || !sameColor(right, q, board)) {
		l = append(l, right)
		right, ok = board.NRight(right, 1)
	}

	// foward left diagonal
	fLDiag, ok := board.NthFowardLeftDiagonal(q.position, 1)
	for ok && (!board.IsOccupied(fLDiag) || !sameColor(fLDiag, q, board)) {
		l = append(l, fLDiag)
		fLDiag, ok = board.NthFowardLeftDiagonal(fLDiag, 1)
	}

	// foward right diagonal
	fRDiag, ok := board.NthFowardRightDiagonal(q.position, 1)
	for ok && (!board.IsOccupied(fRDiag) || !sameColor(fRDiag, q, board)) {
		l = append(l, fRDiag)
		fRDiag, ok = board.NthFowardRightDiagonal(fRDiag, 1)
	}

	// backward left diagonal
	bLDiag, ok := board.NthBackwardLeftDiagonal(q.position, 1)
	for ok && (!board.IsOccupied(bLDiag) || !sameColor(bLDiag, q, board)) {
		l = append(l, bLDiag)
		bLDiag, ok = board.NthBackwardLeftDiagonal(bLDiag, 1)
	}

	// backward right diagonal
	bRDiag, ok := board.NthBackwardRightDiagonal(q.position, 1)
	for ok && (!board.IsOccupied(bRDiag) || !sameColor(bRDiag, q, board)) {
		l = append(l, bLDiag)
		bRDiag, ok = board.NthBackwardRightDiagonal(bRDiag, 1)
	}

	q.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (q *Queen) Move(to utils.Coordinate, board board) bool {
	q.CalculateLegalMoves(board)
	if utils.Contains(q.legalMoves, to) {
		q.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (q Queen) String() string {
	return fmt.Sprintf("%s queen", q.color)
}
