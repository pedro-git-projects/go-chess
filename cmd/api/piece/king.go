package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// check iff king is in any piece list of legal moves

// the King type represents the king piece in a game of chess
type King struct {
	color      Color
	moved      bool
	check      bool
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewKing returns a new king instance
func NewKing(color Color, position utils.Coordinate) *King {
	return &King{
		color:    color,
		moved:    false,
		check:    false,
		position: position,
	}
}

// Color is an accessor for the color field
func (k King) Color() Color {
	return k.color
}

// Position returns the current piece position
func (k King) Position() utils.Coordinate {
	return k.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (k King) LegalMoves() []utils.Coordinate {
	return k.legalMoves
}

// Moved sets the pawn moved field to true
func (k *King) Moved() {
	k.moved = true
}

// SetCheck is a mutator for the check field
func (k *King) SetCheck(c bool) {
	if k.check == c {
		return
	}
	k.check = c
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (k *King) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)
	p0, ok := board.FirstFoward(k.position)
	if ok && !sameColor(p0, k, board) {
		l = append(l, p0)
	}
	p1, ok := board.FirstBackward(k.position)
	if ok && !sameColor(p1, k, board) {
		l = append(l, p1)
	}
	p2, ok := board.NthBackwardLeftDiagonal(k.position, 1)
	if ok && !sameColor(p2, k, board) {
		l = append(l, p2)
	}
	p3, ok := board.NthBackwardRightDiagonal(k.position, 1)
	if ok && !sameColor(p3, k, board) {
		l = append(l, p3)
	}
	p4, ok := board.NthFowardLeftDiagonal(k.position, 1)
	if ok && !sameColor(p4, k, board) {
		l = append(l, p4)
	}
	p5, ok := board.NthFowardRightDiagonal(k.position, 1)
	if ok && !sameColor(p5, k, board) {
		l = append(l, p5)
	}
	p6, ok := board.NLeft(k.position, 1)
	if ok && !sameColor(p6, k, board) {
		l = append(l, p6)
	}
	p7, ok := board.NRight(k.position, 1)
	if ok && !sameColor(p7, k, board) {
		l = append(l, p7)
	}
	k.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (k *King) Move(to utils.Coordinate, board board) bool {
	k.CalculateLegalMoves(board)
	if utils.Contains(k.legalMoves, to) {
		if !k.moved {
			k.moved = true
		}
		k.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (k King) String() string {
	return fmt.Sprintf("%s king", k.color)
}
