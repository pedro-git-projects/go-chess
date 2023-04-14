package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/game"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(table *game.Table) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		// Generate a new client ID
		clientID := utils.GenerateRoomId()

		// Create a new game state and add it to the table
		gameState := game.New()
		table.SetGame(clientID, gameState)

		// Send the client ID to the client
		if err := conn.WriteJSON(map[string]string{"clientID": clientID}); err != nil {
			log.Println(err)
			return
		}

		for {
			// Read a message from the client
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			// Parse the message as a Coordinate
			var coord utils.Coordinate
			if err := json.Unmarshal(message, &coord); err != nil {
				log.Println(err)
				continue
			}

			// Get the game state for the client ID
			gameState := table.Game(clientID)
			if gameState == nil {
				log.Printf("No game state for client ID %s", clientID)
				continue
			}

			// Calculate legal movements for the given coordinate
			legalMoves := gameState.LegalMovements(coord)

			// Marshal the legal movements as JSON and send it to the client
			if err := conn.WriteJSON(legalMoves); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
