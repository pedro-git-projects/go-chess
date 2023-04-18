package piece

import "github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"

func sameColor(postion utils.Coordinate, piece Piece, board board) bool {
	enemy := board.PieceAt(postion)
	if enemy.Color() == piece.Color() {
		return true
	}
	return false
}
