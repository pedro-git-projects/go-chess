// piece defines the piece interface and color enums
package piece

import (
	"github.com/pedro-git-projects/go-chess/src/utils"
)

// Piece is an interface with the restrictions a chess piece
// must satisfy
type Piece interface {
	Color() Color
	Position() utils.Coordinate
	LegalMoves() []utils.Coordinate
	CalculateLegalMoves(board board)
	Move(to utils.Coordinate, board board) bool
}
