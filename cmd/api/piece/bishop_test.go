package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/piece"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestBishopCalculateLegalMoves(t *testing.T) {
	b := board.New()
	bishop := b.PieceAt(utils.NewCoordinate('c', 1))
	bishop.CalculateLegalMoves(b)
	got := bishop.LegalMoves()
	expect := []utils.Coordinate{}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected empty but got %v", got)
	}

	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	bishop.CalculateLegalMoves(b)
	got = bishop.LegalMoves()
	expect = []utils.Coordinate{
		utils.NewCoordinate('d', 2),
		utils.NewCoordinate('e', 3),
		utils.NewCoordinate('f', 4),
		utils.NewCoordinate('g', 5),
		utils.NewCoordinate('h', 6),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestBishopCapture(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	b.MovePiece(utils.NewCoordinate('c', 1), utils.NewCoordinate('d', 2))
	bishop := b.PieceAt(utils.NewCoordinate('d', 2))
	bishop.CalculateLegalMoves(b)
	got := bishop.LegalMoves()
	expect := []utils.Coordinate{
		utils.NewCoordinate('e', 3),
		utils.NewCoordinate('f', 4),
		utils.NewCoordinate('g', 5),
		utils.NewCoordinate('h', 6),
		utils.NewCoordinate('c', 3),
		utils.NewCoordinate('b', 4),
		utils.NewCoordinate('a', 5),
		utils.NewCoordinate('c', 1),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v\n", expect, got)
	}

	b.MovePiece(utils.NewCoordinate('g', 8), utils.NewCoordinate('g', 5))
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('g', 5))
	p := b.PieceAt(utils.NewCoordinate('g', 5))
	if _, ok := p.(*piece.Bishop); !ok {
		t.Errorf("expected bishop, but got %T", reflect.TypeOf(p))
	}

}
