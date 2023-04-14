package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/game"
)

func routes() http.Handler {
	gameTable := game.NewTable()
	router := mux.NewRouter()
	router.HandleFunc("/ws", wsHandler(gameTable))
	return setContentType(enableCORS(router))
}
