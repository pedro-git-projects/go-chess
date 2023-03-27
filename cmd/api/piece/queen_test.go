package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/piece"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestQueenCalculateLegalMovements(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	b.MovePiece(utils.NewCoordinate('d', 7), utils.NewCoordinate('d', 5))
	queen := b.PieceAt(utils.NewCoordinate('d', 1))
	queen.CalculateLegalMoves(b)
	got := queen.LegalMoves()
	want := []utils.Coordinate{
		utils.NewCoordinate('d', 2),
		utils.NewCoordinate('d', 3),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v\n", got, want)
	}
	b.MovePiece(utils.NewCoordinate('d', 1), utils.NewCoordinate('d', 3))
	queen = b.PieceAt(utils.NewCoordinate('d', 3))
	queen.CalculateLegalMoves(b)
	got = queen.LegalMoves()
	want = []utils.Coordinate{
		utils.NewCoordinate('d', 2),
		utils.NewCoordinate('d', 1),
		utils.NewCoordinate('c', 3),
		utils.NewCoordinate('b', 3),
		utils.NewCoordinate('a', 3),
		utils.NewCoordinate('e', 3),
		utils.NewCoordinate('f', 3),
		utils.NewCoordinate('g', 3),
		utils.NewCoordinate('h', 3),
		utils.NewCoordinate('c', 4),
		utils.NewCoordinate('b', 5),
		utils.NewCoordinate('a', 6),
		utils.NewCoordinate('e', 4),
		utils.NewCoordinate('f', 5),
		utils.NewCoordinate('g', 6),
		utils.NewCoordinate('h', 7),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v\n", got, want)
	}
}

func TestQueenCapture(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	b.MovePiece(utils.NewCoordinate('d', 7), utils.NewCoordinate('d', 5))
	b.MovePiece(utils.NewCoordinate('d', 1), utils.NewCoordinate('d', 3))
	ok := b.MovePiece(utils.NewCoordinate('d', 3), utils.NewCoordinate('h', 7))
	got := b.PieceAt(utils.NewCoordinate('h', 7))
	_, isQueen := got.(*piece.Queen)
	if !ok || !isQueen {
		t.Errorf("expected queen but got %T\n", got)
	}
}
