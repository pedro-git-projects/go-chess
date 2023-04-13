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

func boardHandler(game *game.Game) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s := game.MarshalState()
		fmt.Fprintln(w, s)
	}
}

func calculateLegalMovementsHanlder(game *game.Game) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
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

		l := game.LegalMovements(c)
		jsonBytes, err := json.Marshal(l)
		if err != nil {
			fmt.Println("Error marshaling coordinates:", err)
			return
		}
		w.Write(jsonBytes)
	}
}

func movePieceHandler(game *game.Game) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
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

		game.MovePiece(c, c)
		w.WriteHeader(http.StatusOK)
	}
}
