package board

import (
	"fmt"
	"sort"

	"github.com/pedro-git-projects/go-chess/src/piece"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

// Board represents the a chessboard
type Board struct {
	coordinateMatrix [8][8]utils.Coordinate
	state            map[utils.Coordinate]piece.Piece
}

// New returns a pointer to a board
func New() *Board {
	var matrix [8][8]utils.Coordinate
	clc, _ := utils.NewCircularCoordList('a')
	row := 8

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			matrix[i][j] = utils.NewCoordinate(clc.CurrentValue(), row)
			clc.MoveToNext()
		}
		row--
	}

	b := Board{
		coordinateMatrix: matrix,
		state:            make(map[utils.Coordinate]piece.Piece),
	}
	b.initializeBoard()
	return &b
}

// initializeBoard sets up the board state to its default position
func (b *Board) initializeBoard() {
	utils.Copy(b.state, initializeWhitePawns())
	utils.Copy(b.state, initializeBlackPawns())
	utils.Copy(b.state, initializeRooks())
	utils.Copy(b.state, initializeKnights())
	utils.Copy(b.state, initializeBishops())
	utils.Copy(b.state, initializeQueens())
	utils.Copy(b.state, initializeKings())
	utils.Copy(b.state, initilizeBlankSquares())
}

// initilizeWhitePawns sets all white pawns to their default positions
func initializeWhitePawns() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	l, _ := utils.NewCoordList('a')
	for i := 1; i <= 8; i++ {
		m[utils.NewCoordinate(l.CurrentValue(), 2)] = piece.NewPawn(piece.White, utils.NewCoordinate(l.CurrentValue(), 2))
		l.MoveToNext()
	}
	return m
}

// initilizeBlackPawns sets all white pawns to their default positions
func initializeBlackPawns() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	l, _ := utils.NewCoordList('a')
	for i := 1; i <= 8; i++ {
		m[utils.NewCoordinate(l.CurrentValue(), 7)] = piece.NewPawn(piece.Black, utils.NewCoordinate(l.CurrentValue(), 7))
		l.MoveToNext()
	}
	return m
}

// initializeRooks sets all rooks to their starting position
func initializeRooks() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	a1 := utils.NewCoordinate('a', 1)
	h1 := utils.NewCoordinate('h', 1)
	a8 := utils.NewCoordinate('a', 8)
	h8 := utils.NewCoordinate('h', 8)
	m[a1] = piece.NewRook(piece.White, a1)
	m[h1] = piece.NewRook(piece.White, h1)
	m[a8] = piece.NewRook(piece.Black, a8)
	m[h8] = piece.NewRook(piece.Black, h8)
	return m
}

// initializeKnights sets all knights to their starting position
func initializeKnights() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	b1 := utils.NewCoordinate('b', 1)
	g1 := utils.NewCoordinate('g', 1)
	b8 := utils.NewCoordinate('b', 8)
	g8 := utils.NewCoordinate('g', 8)
	m[b1] = piece.NewKnight(piece.White, b1)
	m[g1] = piece.NewKnight(piece.White, g1)
	m[b8] = piece.NewKnight(piece.Black, b8)
	m[g8] = piece.NewKnight(piece.Black, g8)
	return m
}

// initializeBishops sets all knights to their starting position
func initializeBishops() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	c1 := utils.NewCoordinate('c', 1)
	f1 := utils.NewCoordinate('f', 1)
	c8 := utils.NewCoordinate('c', 8)
	f8 := utils.NewCoordinate('f', 8)
	m[c1] = piece.NewBishop(piece.White, c1)
	m[f1] = piece.NewBishop(piece.White, f1)
	m[c8] = piece.NewBishop(piece.Black, c8)
	m[f8] = piece.NewBishop(piece.Black, f8)
	return m
}

// initializeKings sets all kings to their starting position
func initializeKings() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	e1 := utils.NewCoordinate('e', 1)
	e8 := utils.NewCoordinate('e', 8)
	m[e1] = piece.NewKing(piece.White, e1)
	m[e8] = piece.NewKing(piece.Black, e8)
	return m
}

// initializeQueens sets all queens to their starting position
func initializeQueens() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	d1 := utils.NewCoordinate('d', 1)
	d8 := utils.NewCoordinate('d', 8)
	m[d1] = piece.NewQueen(piece.White, d1)
	m[d8] = piece.NewQueen(piece.Black, d8)
	return m
}

// initilizeBlankSquares sets the starting empty squares in the chessboard to empty
func initilizeBlankSquares() map[utils.Coordinate]piece.Piece {
	m := make(map[utils.Coordinate]piece.Piece)
	l, _ := utils.NewCircularCoordList('a')
	row := 3
	for row <= 6 {
		for i := 0; i < 8; i++ {
			m[utils.NewCoordinate(l.CurrentValue(), row)] = piece.Empty{}
			l.MoveToNext()
		}
		row++
	}
	return m
}

// Matrix is an accessor for the board matrix
func (b Board) Matrix() [8][8]utils.Coordinate {
	return b.coordinateMatrix
}

// Find returns true if the coordinate is in the matrix
// false otherwise
func (b Board) Find(c utils.Coordinate) bool {
	for row := 0; row < len(b.coordinateMatrix); row++ {
		for col := 0; col < len(b.coordinateMatrix); col++ {
			if b.coordinateMatrix[row][col].First == c.First &&
				b.coordinateMatrix[row][col].Second == c.Second {
				return true
			}
		}
	}
	return false
}

// FilterByCol returns the column corresponding
// to the letter coordinate passed
func (b Board) FilterByCol(col rune) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)
	for i := 0; i < len(b.coordinateMatrix); i++ {
		for j := 0; j < len(b.coordinateMatrix); j++ {
			if b.coordinateMatrix[i][j].First == col {
				result = append(result, b.coordinateMatrix[i][j])
			}
		}
	}
	return result
}

// FilterByRow returns the column corresponding
// to the number coordinate passed
func (b Board) FilterByRow(row int) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)
	for i := 0; i < len(b.coordinateMatrix); i++ {
		for j := 0; j < len(b.coordinateMatrix); j++ {
			if b.coordinateMatrix[i][j].Second == row {
				result = append(result, b.coordinateMatrix[i][j])
			}
		}
	}
	return result
}

// FowardRightDiagonal returns the foward right diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) FowardRightDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second + 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x++
		y.MoveToNext()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// NLeft returns the nth leftmost postion with respect to a given board coordinate
func (b Board) NLeft(position utils.Coordinate, squares int) (square utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	x := position.Second
	for i := 0; i < squares; i++ {
		if y.Current() == nil {
			return utils.Coordinate{}, false
		}
		y.MoveToPrev()
	}
	nextLeft := utils.NewCoordinate(y.CurrentValue(), x)
	if !b.Find(nextLeft) {
		return utils.Coordinate{}, false
	}
	return nextLeft, true
}

// NRight returns the nth rightmost postion with respect to a given board coordinate
func (b Board) NRight(position utils.Coordinate, squares int) (square utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	x := position.Second
	for i := 0; i < squares; i++ {
		if y.Current() == nil {
			return utils.Coordinate{}, false
		}
		y.MoveToNext()
	}
	nextRight := utils.NewCoordinate(y.CurrentValue(), x)
	if !b.Find(nextRight) {
		return utils.Coordinate{}, false
	}
	return nextRight, true
}

// NthFowardRightDiagonal returns the postion of the nth square in the diagonal right of the current position
func (b Board) NthFowardRightDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second + 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	for i := 0; i < squares-1; i++ {
		if !b.Find(nextDiagonal) {
			return utils.Coordinate{}, false
		}
		x++
		y.MoveToNext()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	return nextDiagonal, true
}

// FowardLeftDiagonal returns the foward left diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) FowardLeftDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second + 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x++
		y.MoveToPrev()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// NthFowardLeftDiagonal returns the postion of the nth square in the diagonal right of the current position
func (b Board) NthFowardLeftDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second + 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	for i := 0; i < squares-1; i++ {
		if !b.Find(nextDiagonal) {
			return utils.Coordinate{}, false
		}
		x++
		y.MoveToPrev()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	return nextDiagonal, true
}

// BackwardLeftDiagonal returns the backward left diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) BackwardLeftDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second - 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x--
		y.MoveToPrev()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// NthBackwardLeftDiagonal returns the postion behind of the nth square in the diagonal left of the current position
func (b Board) NthBackwardLeftDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second - 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	for i := 0; i < squares-1; i++ {
		if !b.Find(nextDiagonal) {
			return utils.Coordinate{}, false
		}
		x--
		y.MoveToPrev()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	return nextDiagonal, true
}

// BackwardRightDiagonal  the backward right diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) BackwardRightDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second - 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x--
		y.MoveToNext()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// NthBackwardRightDiagonal returns the postion behind of the nth square in the diagonal right of the current position
func (b Board) NthBackwardRightDiagonal(position utils.Coordinate, squares int) (diag utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second - 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	for i := 0; i < squares-1; i++ {
		if !b.Find(nextDiagonal) {
			return utils.Coordinate{}, false
		}
		x--
		y.MoveToNext()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	if !b.Find(nextDiagonal) {
		return utils.Coordinate{}, false
	}

	return nextDiagonal, true
}

// FirstFoward returns the next square in front of the current one as well as a true value if it succeeds
// it returns an empty coordinate and not ok else
func (b Board) FirstFoward(position utils.Coordinate) (foward utils.Coordinate, ok bool) {
	next := position.Second + 1
	if next <= 8 {
		return utils.NewCoordinate(position.First, next), true
	}
	return utils.Coordinate{}, false
}

// FirstBackward returns the next square behind the current piece if it is white, the first foward if it is black
// it returns an empty coordinate object and a false flag if it fails.
func (b Board) FirstBackward(position utils.Coordinate) (backward utils.Coordinate, ok bool) {
	next := position.Second - 1
	if next <= 8 && next >= 1 {
		return utils.NewCoordinate(position.First, next), true
	}
	return utils.Coordinate{}, false

}

// NFoward returns the nth square in front of current as well as a true value if it succeeds
// it returns an empty coordinate and not ok else
func (b Board) NFoward(position utils.Coordinate, squares int) (foward utils.Coordinate, ok bool) {
	next := position.Second + squares
	if next <= 8 && next >= 1 {
		return utils.NewCoordinate(position.First, next), true
	}
	return utils.Coordinate{}, false
}

// NBackward returns the nth square in front of current if the pieces are black, behind if they're white
// as well as a true value if it succeeds
// it returns an empty coordinate and not ok else
func (b Board) NBackward(position utils.Coordinate, squares int) (backward utils.Coordinate, ok bool) {
	next := position.Second - squares
	if next <= 8 && next >= 1 {
		return utils.NewCoordinate(position.First, next), true
	}
	return utils.Coordinate{}, false
}

// FowardRightL returns the position two squares foward one right from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) FowardRightL(position utils.Coordinate) (frl utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsNextNil() {
		return
	}
	y.MoveToNext()
	if y.Current() == nil {
		return
	}
	x := position.Second + 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// FowardLeftL returns the position two squares foward one left from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) FowardLeftL(position utils.Coordinate) (fll utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsPrevNil() {
		return
	}
	y.MoveToPrev()
	if y.Current() == nil {
		return
	}
	x := position.Second + 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)
	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// BackwardRightL returns the position two squares backward one right from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) BackwardRightL(position utils.Coordinate) (brl utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsNextNil() {
		return
	}
	y.MoveToNext()
	if y.Current() == nil {
		return
	}
	x := position.Second - 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// BackwardLeftL returns the position two squares backward one left from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) BackwardLeftL(position utils.Coordinate) (brl utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsPrevNil() {
		return
	}
	y.MoveToPrev()
	if y.Current() == nil {
		return
	}
	x := position.Second - 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// RightDownLateralL returns the position two squares right one down from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) RightDownLateralL(position utils.Coordinate) (rdll utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsNextNil() {
		return
	}
	y.MoveToNext()
	if y.IsNextNil() {
		return
	}
	y.MoveToNext()
	if y.Current() == nil {
		return
	}
	x := position.Second - 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// LeftDownLateralL returns the position two squares left one down from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) LeftDownLateralL(position utils.Coordinate) (ldll utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsPrevNil() {
		return
	}
	y.MoveToPrev()
	if y.IsPrevNil() {
		return
	}
	y.MoveToPrev()
	if y.Current() == nil {
		return
	}
	x := position.Second - 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// LeftUpLateralL returns the position two squares left one up from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) LeftUpLateralL(position utils.Coordinate) (lull utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	if y.IsPrevNil() {
		return
	}
	y.MoveToPrev()
	if y.IsPrevNil() {
		return
	}
	y.MoveToPrev()
	if y.Current() == nil {
		return
	}
	x := position.Second + 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// RightUpLateralL returns the position two squares right one up from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) RightUpLateralL(position utils.Coordinate) (rull utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)

	if y.IsNextNil() {
		return
	}
	y.MoveToNext()
	if y.IsNextNil() {
		return
	}
	y.MoveToNext()
	if y.Current() == nil {
		return
	}
	x := position.Second + 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// CalcAllLs calculates all L patterns from a given position
// and returns the corresponding possibly empty coordinate slice
func (b Board) CalcAllLs(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	c, ok := b.FowardRightL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.FowardLeftL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.BackwardLeftL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.BackwardRightL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.RightDownLateralL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.LeftDownLateralL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.LeftUpLateralL(position)
	if ok {
		result = append(result, c)
	}
	c, ok = b.RightUpLateralL(position)
	if ok {
		result = append(result, c)
	}
	return result
}

// State returns the board map
func (b Board) State() map[utils.Coordinate]piece.Piece {
	return b.state
}

// TODO: ROTATE 90 degrees
// StateStr retuns a string representation of the chessboard
// in the form of [coord piece]
func (b Board) StateStr() string {
	s := ""
	keys := []string{}
	for k := range b.state {
		keys = append(keys, string(k.First)+fmt.Sprint(k.Second))
	}
	sort.Strings(keys)
	counter := 1
	for _, k := range keys {
		coord, err := utils.CoordFromStr(k)
		if err != nil {
			fmt.Println(err)
		}
		if counter%8 == 0 {
			s += fmt.Sprintf("[%v %v]\n", k, b.state[coord])
			counter++
		} else {
			s += fmt.Sprintf("[%v %v]", k, b.state[coord])
			counter++
		}
	}
	return s
}

// Marshal is a function that takes the board state
// and marshals it into a json string
// TODO: separate piece type and color (maybe)
func (b Board) Marshal() string {
	result := ""
	delimiters := "["
	keys := []string{}
	for k := range b.state {
		keys = append(keys, string(k.First)+fmt.Sprint(k.Second))
	}
	sort.Strings(keys)
	result += delimiters
	for i, k := range keys {
		coord, err := utils.CoordFromStr(k)
		if err != nil {
			fmt.Println(err)
		}
		if i != len(keys)-1 {
			result += fmt.Sprintf("{%q: %q, %q: %q},", "coordinate", k, "piece", b.state[coord])

		} else {
			result += fmt.Sprintf("{%q: %q, %q: %q}", "coordinate", k, "piece", b.state[coord])
		}
	}
	delimiters = "]"
	result += delimiters
	return result
}

// TODO
func (b Board) UnmarshalCoordinate(json string, err error) {

}

// IsOccupied returns true if the passed coordinate is
// not nil, false otherwise
func (b Board) IsOccupied(c utils.Coordinate) bool {
	_, ok := b.state[c].(piece.Empty)
	if b.state[c] == nil || ok {
		return false
	}
	return true
}

// PieceAt returns the piece at a given coordinate
func (b Board) PieceAt(c utils.Coordinate) piece.Piece {
	return b.state[c]
}

// MovePiece moves clears the current position of a piece in the map
// and associates the piece as the value of the new coordinate key
// if the coordinate exist
func (b *Board) MovePiece(from utils.Coordinate, to utils.Coordinate) (ok bool) {
	_, ok = b.state[from] // checks if the coordinate has a piece
	if !ok {
		return ok
	}
	p := b.PieceAt(from) // gets the piece
	ok = p.Move(to, b)   // updates piece internal state
	if ok {              // updates game state
		b.state[from] = piece.Empty{} // clears starting position
		b.state[to] = p               // updates final boardstate postion
		return ok
	}
	return ok // an invalid move was attempted
}
