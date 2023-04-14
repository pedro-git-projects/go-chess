package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/game"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/utils"
)

type Coordinate struct {
	Coordinate string `json:"coordinate"`
}

// boardHandler is a closure on a httprouter.Handle which
// tries to fetch a game object from the table.
// If it fails, a new entry is created using the clientID.
func boardHandler(table *game.Table) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		clientID := p.ByName("clientID")
		gameState := table.Game(clientID)
		if gameState == nil {
			gameState = game.New()
			table.SetGame(clientID, gameState)
		}
		s := gameState.MarshalState()
		fmt.Fprintln(w, s)
	}
}

func calculateLegalMovementsHanlder(table *game.Table) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		clientID := p.ByName("clientID")
		gameState := table.Game(clientID)
		if gameState == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var coord Coordinate
		err := json.NewDecoder(r.Body).Decode(&coord)
		if err != nil {
			log.Println("Error decoding JSON request body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c, err := utils.CoordFromStr(coord.Coordinate)
		if err != nil {
			log.Println("Error converting object to coordinate:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		l := gameState.LegalMovements(c)
		jsonBytes, err := json.Marshal(l)
		if err != nil {
			fmt.Println("Error marshaling coordinates:", err)
			return
		}
		w.Write(jsonBytes)
	}
}

// TODO get final coordinate from request body
func movePieceHandler(table *game.Table) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		clientID := p.ByName("clientID")
		gameState := table.Game(clientID)
		if gameState == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var coord Coordinate
		err := json.NewDecoder(r.Body).Decode(&coord)
		if err != nil {
			log.Println("Error decoding JSON request body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c, err := utils.CoordFromStr(coord.Coordinate)
		if err != nil {
			log.Println("Error converting object to coordinate:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		gameState.MovePiece(c, c)
		w.WriteHeader(http.StatusOK)
	}
}
