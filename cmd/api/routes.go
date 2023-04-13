package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pedro-git-projects/projeto-integrado-frontend/cmd/api/game"
)

func routes() http.Handler {
	gameState := game.New()
	router := httprouter.New()
	router.GET("/board", boardHandler(gameState))
	router.POST("/move", movePieceHandler(gameState))
	router.POST("/calc", calculateLegalMovementsHanlder(gameState))
	return setContentType(enableCORS(router))
}
