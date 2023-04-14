package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/game"
)

func routes() http.Handler {
	gameTable := game.NewTable()
	router := httprouter.New()
	router.GET("/room/new", newRoomHandler(gameTable))
	router.GET("/board/:clientID", boardHandler(gameTable))
	router.POST("/move/:clientID", movePieceHandler(gameTable))
	router.POST("/calc/:clientID", calculateLegalMovementsHanlder(gameTable))
	return setContentType(enableCORS(router))
}
