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

func (t *Table) Game(clientID string) *Game {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.games[clientID]
}

func (t *Table) SetGame(clientID string, gameState *Game) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.games[clientID] = gameState
}
