package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/piece"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestKingCalculateLegalMoves(t *testing.T) {
	b := board.New()
	king := b.PieceAt(utils.NewCoordinate('e', 1))
	king.CalculateLegalMoves(b)
	got := king.LegalMoves()
	expect := []utils.Coordinate{}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}

	b.MovePiece(utils.NewCoordinate('f', 2), utils.NewCoordinate('f', 4))
	king.CalculateLegalMoves(b)
	got = king.LegalMoves()
	expect = []utils.Coordinate{
		utils.NewCoordinate('f', 2),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestKingMovement(t *testing.T) {
	b := board.New()

	// white king
	b.MovePiece(utils.NewCoordinate('f', 2), utils.NewCoordinate('f', 4))
	ok := b.MovePiece(utils.NewCoordinate('e', 1), utils.NewCoordinate('f', 2))
	if !ok {
		t.Error("Expected true but got false")
	}
	got := b.PieceAt(utils.NewCoordinate('f', 2))
	if _, ok := got.(*piece.King); !ok {
		t.Errorf("Expected king but got %v", got)
	}

	// black king
	b.MovePiece(utils.NewCoordinate('f', 7), utils.NewCoordinate('f', 5))
	ok = b.MovePiece(utils.NewCoordinate('e', 8), utils.NewCoordinate('f', 7))
	if !ok {
		t.Error("Expected true but got false")
	}
	got = b.PieceAt(utils.NewCoordinate('f', 7))
	if _, ok := got.(*piece.King); !ok {
		t.Errorf("Expected king but got %v", got)
	}

}

func TestKingCapture(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('f', 2), utils.NewCoordinate('f', 4)) // white
	b.MovePiece(utils.NewCoordinate('e', 7), utils.NewCoordinate('e', 5)) // black
	b.MovePiece(utils.NewCoordinate('e', 1), utils.NewCoordinate('f', 2)) // white
	b.MovePiece(utils.NewCoordinate('e', 5), utils.NewCoordinate('e', 4)) // black
	b.MovePiece(utils.NewCoordinate('f', 2), utils.NewCoordinate('e', 3)) // white
	b.MovePiece(utils.NewCoordinate('e', 3), utils.NewCoordinate('e', 4)) // white king captures black pawn
	got := b.PieceAt(utils.NewCoordinate('e', 4))
	if _, ok := got.(*piece.King); !ok {
		t.Errorf("Expected king but got %v", got)
	}
}
