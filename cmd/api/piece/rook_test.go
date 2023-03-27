package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/piece"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestRookCalculateLegalMoves(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('a', 2), utils.NewCoordinate('a', 4))
	rook := b.PieceAt(utils.NewCoordinate('a', 1))
	rook.CalculateLegalMoves(b)
	got := rook.LegalMoves()
	expect := []utils.Coordinate{
		utils.NewCoordinate('a', 2),
		utils.NewCoordinate('a', 3),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected %v but got %v", expect, got)
	}

	b.MovePiece(utils.NewCoordinate('a', 1), utils.NewCoordinate('a', 3))
	rook.CalculateLegalMoves(b)
	got = rook.LegalMoves()
	expect = []utils.Coordinate{
		utils.NewCoordinate('b', 3),
		utils.NewCoordinate('c', 3),
		utils.NewCoordinate('d', 3),
		utils.NewCoordinate('e', 3),
		utils.NewCoordinate('f', 3),
		utils.NewCoordinate('g', 3),
		utils.NewCoordinate('h', 3),
		utils.NewCoordinate('a', 2),
		utils.NewCoordinate('a', 1),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected %v but got %v", expect, got)
	}
}

func TestWhiteRookCaputre(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('a', 2), utils.NewCoordinate('a', 4))
	b.MovePiece(utils.NewCoordinate('a', 1), utils.NewCoordinate('a', 3))
	b.MovePiece(utils.NewCoordinate('a', 3), utils.NewCoordinate('c', 3))
	b.MovePiece(utils.NewCoordinate('c', 3), utils.NewCoordinate('c', 7))
	got := b.PieceAt(utils.NewCoordinate('c', 7))
	if _, ok := got.(*piece.Rook); !ok {
		t.Errorf("Expected rook but got %v", got)
	}
}

func TestBlackRookCapture(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('a', 7), utils.NewCoordinate('a', 5))
	b.MovePiece(utils.NewCoordinate('a', 8), utils.NewCoordinate('a', 6))
	b.MovePiece(utils.NewCoordinate('a', 6), utils.NewCoordinate('c', 6))
	b.MovePiece(utils.NewCoordinate('c', 6), utils.NewCoordinate('c', 2))
	got := b.PieceAt(utils.NewCoordinate('c', 2))
	if _, ok := got.(*piece.Rook); !ok {
		t.Errorf("Expected rook but got %v", got)
	}
}
