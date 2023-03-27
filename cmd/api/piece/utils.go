package piece

import (
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func sameColor(postion utils.Coordinate, piece Piece, board board) bool {
	enemy := board.PieceAt(postion)
	if enemy.Color() == piece.Color() {
		return true
	}
	return false
}
