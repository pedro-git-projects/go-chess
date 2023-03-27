package piece

import "github.com/pedro-git-projects/go-chess/src/utils"

// Interface to break dependency cycle
type board interface {
	Find(c utils.Coordinate) bool
	PieceAt(c utils.Coordinate) Piece
	FilterByCol(col rune) []utils.Coordinate
	FilterByRow(row int) []utils.Coordinate
	FowardRightDiagonal(position utils.Coordinate) []utils.Coordinate
	NLeft(position utils.Coordinate, squares int) (square utils.Coordinate, ok bool)
	NRight(position utils.Coordinate, squares int) (square utils.Coordinate, ok bool)
	NthFowardRightDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool)
	FowardLeftDiagonal(position utils.Coordinate) []utils.Coordinate
	NthFowardLeftDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool)
	BackwardLeftDiagonal(position utils.Coordinate) []utils.Coordinate
	NthBackwardLeftDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool)
	BackwardRightDiagonal(position utils.Coordinate) []utils.Coordinate
	NthBackwardRightDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool)
	FirstFoward(position utils.Coordinate) (foward utils.Coordinate, ok bool)
	FirstBackward(position utils.Coordinate) (backward utils.Coordinate, ok bool)
	NFoward(position utils.Coordinate, squares int) (foward utils.Coordinate, ok bool)
	NBackward(position utils.Coordinate, squares int) (backward utils.Coordinate, ok bool)
	FowardRightL(position utils.Coordinate) (frl utils.Coordinate, ok bool)
	FowardLeftL(position utils.Coordinate) (fll utils.Coordinate, ok bool)
	BackwardRightL(position utils.Coordinate) (brl utils.Coordinate, ok bool)
	BackwardLeftL(position utils.Coordinate) (brl utils.Coordinate, ok bool)
	RightDownLateralL(position utils.Coordinate) (rdll utils.Coordinate, ok bool)
	LeftDownLateralL(position utils.Coordinate) (ldll utils.Coordinate, ok bool)
	LeftUpLateralL(position utils.Coordinate) (lull utils.Coordinate, ok bool)
	RightUpLateralL(position utils.Coordinate) (rull utils.Coordinate, ok bool)
	CalcAllLs(position utils.Coordinate) []utils.Coordinate
	State() map[utils.Coordinate]Piece
	IsOccupied(c utils.Coordinate) bool
	MovePiece(from utils.Coordinate, to utils.Coordinate) (ok bool)
}
