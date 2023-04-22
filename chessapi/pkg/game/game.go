package game

import (
	"errors"
	"fmt"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/board"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/piece"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
)

type Game struct {
	board       *board.Board
	currentTurn piece.Color
	client1     *Client
	client2     *Client
}

// New returns a pointer to a Game
// the zero values are fully usable.
func NewGame() *Game {
	return &Game{
		board:       board.New(),
		currentTurn: piece.White,
		client1:     nil,
		client2:     nil,
	}
}

// AddClient tries to add a Client to a game
// it returns an error if it fails
func (g *Game) AddClient(c *Client) error {
	if g.client1 == nil {
		g.client1 = c
		c.color = piece.White
		return nil
	}
	if g.client2 == nil {
		g.client2 = c
		c.color = piece.Black
		return nil
	}
	return errors.New("Game is full")
}

func (game *Game) PieceColor(c utils.Coordinate) piece.Color {
	return game.board.PieceAt(c).Color()
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
func (game *Game) MovePiece(from, to utils.Coordinate, clientColor piece.Color) {
	p := game.board.PieceAt(from)
	if p.Color() != game.currentTurn || clientColor != game.currentTurn {
		return
	}
	ok := game.board.MovePiece(from, to)
	if ok {
		game.setCurrentTurn(p.Color())
	}
}

func (game *Game) LegalMovements(c utils.Coordinate, clientColor piece.Color) []utils.Coordinate {
	if clientColor == game.currentTurn {
		p := game.board.PieceAt(c)
		p.CalculateLegalMoves(game.board)
		return p.LegalMoves()

	} else {
		return []utils.Coordinate{}
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
