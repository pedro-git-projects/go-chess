package board_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestFind(t *testing.T) {
	board := board.New()
	got := board.Find(utils.NewCoordinate('a', 3))
	expect := true

	if got != expect {
		t.Errorf("expected %t but got %t", expect, got)
	}

	got = board.Find(utils.NewCoordinate('z', 1))
	expect = false

	if got != expect {
		t.Errorf("expected %t but got %t", expect, got)
	}
}

func TestFilterByCol(t *testing.T) {
	board := board.New()
	got := board.FilterByCol('a')
	expect := []utils.Coordinate{
		utils.NewCoordinate('a', 8),
		utils.NewCoordinate('a', 7),
		utils.NewCoordinate('a', 6),
		utils.NewCoordinate('a', 5),
		utils.NewCoordinate('a', 4),
		utils.NewCoordinate('a', 3),
		utils.NewCoordinate('a', 2),
		utils.NewCoordinate('a', 1),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestFilterByRow(t *testing.T) {
	board := board.New()
	got := board.FilterByRow(1)
	expect := []utils.Coordinate{
		utils.NewCoordinate('a', 1),
		utils.NewCoordinate('b', 1),
		utils.NewCoordinate('c', 1),
		utils.NewCoordinate('d', 1),
		utils.NewCoordinate('e', 1),
		utils.NewCoordinate('f', 1),
		utils.NewCoordinate('g', 1),
		utils.NewCoordinate('h', 1),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestFowardRightDiagonal(t *testing.T) {
	board := board.New()
	got := board.FowardRightDiagonal(utils.NewCoordinate('d', 1))
	expect := []utils.Coordinate{
		utils.NewCoordinate('e', 2),
		utils.NewCoordinate('f', 3),
		utils.NewCoordinate('g', 4),
		utils.NewCoordinate('h', 5),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestFowardLeftDiagonal(t *testing.T) {
	board := board.New()
	got := board.FowardLeftDiagonal(utils.NewCoordinate('f', 5))
	expect := []utils.Coordinate{
		utils.NewCoordinate('e', 6),
		utils.NewCoordinate('d', 7),
		utils.NewCoordinate('c', 8),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestBackwardLeftDiagonal(t *testing.T) {
	board := board.New()
	got := board.BackwardLeftDiagonal(utils.NewCoordinate('e', 4))
	expect := []utils.Coordinate{
		utils.NewCoordinate('d', 3),
		utils.NewCoordinate('c', 2),
		utils.NewCoordinate('b', 1),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestBackwardRightDiagonal(t *testing.T) {
	board := board.New()
	got := board.BackwardRightDiagonal(utils.NewCoordinate('e', 4))
	expect := []utils.Coordinate{
		utils.NewCoordinate('f', 3),
		utils.NewCoordinate('g', 2),
		utils.NewCoordinate('h', 1),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestFirstFoward(t *testing.T) {
	board := board.New()
	got, ok := board.FirstFoward(utils.NewCoordinate('e', 4))
	expect := utils.NewCoordinate('e', 5)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.FirstFoward(utils.NewCoordinate('e', 8))
	if ok {
		t.Errorf("expected %t but got %t", true, ok)
	}
}

func TestNFoward(t *testing.T) {
	board := board.New()
	got, ok := board.NFoward(utils.NewCoordinate('e', 4), 4)
	expect := utils.NewCoordinate('e', 8)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.NFoward(utils.NewCoordinate('e', 4), 5)
	if ok {
		t.Errorf("expected %t but got %t", true, ok)
	}
}

func TestFowardRightL(t *testing.T) {
	board := board.New()
	got, ok := board.FowardRightL(utils.NewCoordinate('f', 1))

	expect := utils.NewCoordinate('g', 3)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.FowardRightL(utils.NewCoordinate('h', 3))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestFowardLeftL(t *testing.T) {
	board := board.New()
	got, ok := board.FowardLeftL(utils.NewCoordinate('f', 1))

	expect := utils.NewCoordinate('e', 3)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.FowardLeftL(utils.NewCoordinate('a', 3))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestBackwardRightL(t *testing.T) {
	board := board.New()
	got, ok := board.BackwardRightL(utils.NewCoordinate('e', 3))

	expect := utils.NewCoordinate('f', 1)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.BackwardRightL(utils.NewCoordinate('f', 2))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestBackwardLeftL(t *testing.T) {
	board := board.New()
	got, ok := board.BackwardLeftL(utils.NewCoordinate('e', 3))

	expect := utils.NewCoordinate('d', 1)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.BackwardLeftL(utils.NewCoordinate('f', 2))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestRightDownLateralL(t *testing.T) {
	board := board.New()
	got, ok := board.RightDownLateralL(utils.NewCoordinate('e', 3))

	expect := utils.NewCoordinate('g', 2)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.RightDownLateralL(utils.NewCoordinate('g', 3))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestLeftDownLateralL(t *testing.T) {
	board := board.New()
	got, ok := board.LeftDownLateralL(utils.NewCoordinate('e', 3))

	expect := utils.NewCoordinate('c', 2)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.LeftDownLateralL(utils.NewCoordinate('b', 3))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestLeftUpLateralL(t *testing.T) {
	board := board.New()
	got, ok := board.LeftUpLateralL(utils.NewCoordinate('e', 3))

	expect := utils.NewCoordinate('c', 4)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.LeftUpLateralL(utils.NewCoordinate('b', 3))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestRightpLateralL(t *testing.T) {
	board := board.New()
	got, ok := board.RightUpLateralL(utils.NewCoordinate('e', 3))

	expect := utils.NewCoordinate('g', 4)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.RightUpLateralL(utils.NewCoordinate('g', 3))
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestCalcAllLs(t *testing.T) {
	board := board.New()
	got := board.CalcAllLs(utils.NewCoordinate('d', 4))

	expect := []utils.Coordinate{
		utils.NewCoordinate('e', 6),
		utils.NewCoordinate('c', 6),
		utils.NewCoordinate('c', 2),
		utils.NewCoordinate('e', 2),
		utils.NewCoordinate('f', 3),
		utils.NewCoordinate('b', 3),
		utils.NewCoordinate('b', 5),
		utils.NewCoordinate('f', 5),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestNthFowardRightDiagonal(t *testing.T) {
	board := board.New()
	got, ok := board.NthFowardRightDiagonal(utils.NewCoordinate('e', 2), 1)

	expect := utils.NewCoordinate('f', 3)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.NthFowardRightDiagonal(utils.NewCoordinate('e', 2), 4)
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestNthFowardLeftDiagonal(t *testing.T) {
	board := board.New()
	got, ok := board.NthFowardLeftDiagonal(utils.NewCoordinate('e', 2), 3)

	expect := utils.NewCoordinate('b', 5)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.NthFowardRightDiagonal(utils.NewCoordinate('e', 2), 5)
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestNthBackwardLeftDiagonal(t *testing.T) {
	board := board.New()
	got, ok := board.NthBackwardLeftDiagonal(utils.NewCoordinate('e', 2), 1)

	expect := utils.NewCoordinate('d', 1)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.NthBackwardRightDiagonal(utils.NewCoordinate('e', 2), 2)
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestNthBackwardRightDiagonal(t *testing.T) {
	board := board.New()
	got, ok := board.NthBackwardRightDiagonal(utils.NewCoordinate('e', 2), 1)

	expect := utils.NewCoordinate('f', 1)

	if !reflect.DeepEqual(got, expect) || !ok {
		t.Errorf("expected %v but got %v", expect, got)
	}

	got, ok = board.NthBackwardRightDiagonal(utils.NewCoordinate('e', 2), 2)
	if ok {
		t.Errorf("expected %t but got %t", false, ok)
	}
}

func TestNLeft(t *testing.T) {
	board := board.New()
	square, ok := board.NLeft(utils.NewCoordinate('c', 1), 2)
	if !ok {
		t.Error("Expected true but got false")
	}
	expect := utils.NewCoordinate('a', 1)
	if !reflect.DeepEqual(square, expect) {
		t.Errorf("Expected %v but got %v", expect, square)
	}

	square, ok = board.NLeft(utils.NewCoordinate('a', 1), 2)
	if ok {
		t.Error("Expected false but got true")
	}
}

func TestNRight(t *testing.T) {
	board := board.New()
	square, ok := board.NRight(utils.NewCoordinate('c', 1), 2)
	if !ok {
		t.Error("Expected true but got false")
	}
	expect := utils.NewCoordinate('e', 1)
	if !reflect.DeepEqual(square, expect) {
		t.Errorf("Expected %v but got %v", expect, square)
	}

	square, ok = board.NRight(utils.NewCoordinate('g', 1), 2)
	if ok {
		t.Error("Expected false but got true")
	}
}
