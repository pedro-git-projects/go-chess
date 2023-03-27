package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestMovement(t *testing.T) {
	b := board.New()
	p1 := b.PieceAt(utils.NewCoordinate('a', 4))
	b.MovePiece(utils.NewCoordinate('a', 2), utils.NewCoordinate('a', 4))
	p2 := b.PieceAt(utils.NewCoordinate('a', 4))
	if reflect.DeepEqual(p1, p2) {
		t.Errorf("expected white pawn, got %s", p2)
	}
}

func TestPawnCalculateLegalMoves(t *testing.T) {
	b := board.New()
	pawn := b.PieceAt(utils.NewCoordinate('a', 2))
	pawn.CalculateLegalMoves(b)
	got := pawn.LegalMoves()

	expect := []utils.Coordinate{
		utils.NewCoordinate('a', 3),
		utils.NewCoordinate('a', 4),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestCapturePiece(t *testing.T) {
	b := board.New()

	b.MovePiece(utils.NewCoordinate('a', 2), utils.NewCoordinate('a', 4))
	b.MovePiece(utils.NewCoordinate('b', 7), utils.NewCoordinate('b', 5))
	p1 := b.PieceAt(utils.NewCoordinate('b', 5))
	b.MovePiece(utils.NewCoordinate('a', 4), utils.NewCoordinate('b', 5))
	p2 := b.PieceAt(utils.NewCoordinate('b', 5))

	if reflect.DeepEqual(p1, p2) {
		t.Errorf("expected different pieces, but got %v and %v", p1, p2)
	}
}

func TestInvalidPostion(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('a', 2), utils.NewCoordinate('a', 4))
	ok := b.MovePiece(utils.NewCoordinate('a', 4), utils.NewCoordinate('a', 2))
	if ok {
		t.Error("Expected false, but got true")
	}
	ok = b.MovePiece(utils.NewCoordinate('a', 4), utils.NewCoordinate('b', 5))
	if ok {
		t.Error("Expected false, but got true")
	}
}

func TestCaptureSameColor(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('c', 2), utils.NewCoordinate('c', 3))
	ok := b.MovePiece(utils.NewCoordinate('b', 2), utils.NewCoordinate('c', 3))
	if ok {
		t.Error("Self capture")
	}
}
