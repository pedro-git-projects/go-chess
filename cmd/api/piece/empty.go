package piece

import "github.com/pedro-git-projects/go-chess/src/utils"

// Empty represents an empty space
// it exists solely for the purpose of
// avoiding nil pointer derefrences when trying to perform
// a move operation on an empty board position
type Empty struct {
	color      Color
	moved      bool
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

func (Empty) Color() Color                               { return None }
func (Empty) Position() utils.Coordinate                 { return utils.Coordinate{} }
func (Empty) LegalMoves() []utils.Coordinate             { return make([]utils.Coordinate, 0) }
func (Empty) CalculateLegalMoves(board board)            {}
func (Empty) Move(to utils.Coordinate, board board) bool { return false }
func (e Empty) String() string                           { return "empty" }
