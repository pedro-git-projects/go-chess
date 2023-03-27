package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Pawn type represents pawn pieces in a game of chess
type Pawn struct {
	color      Color
	moved      bool
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewPawn returns a new pawn instance
func NewPawn(color Color, position utils.Coordinate) *Pawn {
	return &Pawn{
		color:    color,
		moved:    false,
		position: position,
	}
}

// Color is an accessor for the color field
func (p Pawn) Color() Color {
	return p.color
}

// Position returns the current piece position
func (p Pawn) Position() utils.Coordinate {
	return p.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (p Pawn) LegalMoves() []utils.Coordinate {
	return p.legalMoves
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (p *Pawn) CalculateLegalMoves(board board) {
	r := make([]utils.Coordinate, 0)
	if p.color == White {
		// walk foward
		position, ok := board.FirstFoward(p.position)
		if !board.IsOccupied(position) && ok {
			r = append(r, position)
		}
		if !ok {
			return
		}

		// piece capture
		lDiagonal, ok := board.NthFowardRightDiagonal(p.position, 1)
		if board.IsOccupied(lDiagonal) && ok && !sameColor(lDiagonal, p, board) {
			r = append(r, lDiagonal)

		}

		// piece capture
		rDiagonal, ok := board.NthBackwardLeftDiagonal(p.position, 1)
		if board.IsOccupied(rDiagonal) && ok && !sameColor(rDiagonal, p, board) {
			r = append(r, rDiagonal)
		}

		// move two squares if it is the first movement
		if !p.moved {
			position, ok := board.NFoward(p.position, 2)
			if ok && !board.IsOccupied(position) {
				r = append(r, position)
			}
		}
	}

	if p.color == Black {
		// walk foward
		position, ok := board.FirstBackward(p.position)
		if !board.IsOccupied(position) && ok {
			r = append(r, position)
		}
		if !ok {
			return
		}

		// piece capture
		lDiagonal, ok := board.NthBackwardRightDiagonal(p.position, 1)
		if board.IsOccupied(lDiagonal) && ok && !sameColor(lDiagonal, p, board) {
			r = append(r, lDiagonal)

		}

		// piece capture
		rDiagonal, ok := board.NthBackwardLeftDiagonal(p.position, 1)
		if board.IsOccupied(rDiagonal) && ok && !sameColor(rDiagonal, p, board) {
			r = append(r, rDiagonal)
		}

		// move two squares if it is the first movement
		if !p.moved {
			position, ok := board.NBackward(p.position, 2)
			if ok && !board.IsOccupied(position) {
				r = append(r, position)
			}
		}
	}

	p.legalMoves = r
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (p *Pawn) Move(to utils.Coordinate, board board) bool {
	p.CalculateLegalMoves(board)
	if utils.Contains(p.legalMoves, to) {
		p.position = to
		if !p.moved {
			p.moved = true
		}
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (p Pawn) String() string {
	return fmt.Sprintf("%s pawn", p.color)
}
