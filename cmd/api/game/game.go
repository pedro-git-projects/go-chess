package game

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/piece"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

type Game struct {
	board       *board.Board
	currentTurn piece.Color
}

// New returns a pointer to a Game
// the zero values are fully usable.
func New() *Game {
	return &Game{
		board:       board.New(),
		currentTurn: piece.White,
	}
}

// setCurrentTurn takes a piece color and sets the currentTurn color
// to the opposite color
func (game *Game) setCurrentTurn(color piece.Color) {
	if color == piece.White {
		game.currentTurn = piece.Black
		return
	}
	game.currentTurn = piece.White
}

// MovePiece takes two coordiantes, from and to
// checks if the checks if the current turns is of adequate color
// moves the piece and updates the current color turn if it is fit
func (game *Game) MovePiece(from, to utils.Coordinate) {
	p := game.board.PieceAt(from)
	if p.Color() != game.currentTurn {
		return
	}
	ok := game.board.MovePiece(from, to)
	if ok {
		game.setCurrentTurn(p.Color())
	}
}

// PrintBoard prints the current board state to the os.Stdout
func (g Game) PrintBoard() {
	fmt.Println(g.board.StateStr())
}

// MarshalState returns the current board state as a JSON object
func (g Game) MarshalState() string {
	return g.board.Marshal()
}
