package game

import (
	"sync"
)

// type Table represents an associative array
// of games and clientIDs
type Table struct {
	mu    sync.Mutex
	games map[string]*Game
}

func NewTable() *Table {
	return &Table{
		games: make(map[string]*Game),
	}
}

func (t *Table) Game(roomID string) *Game {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.games[roomID]
}

func (t *Table) SetGame(roomID string, gameState *Game) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.games[roomID] = gameState
}
