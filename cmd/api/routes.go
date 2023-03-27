package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()
	router.Handler(http.MethodGet, "/board", http.HandlerFunc(boardHandler))
	return setContentType(enableCORS(router))
}
