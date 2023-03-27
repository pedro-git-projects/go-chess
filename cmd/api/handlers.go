package main

import (
	"fmt"
	"net/http"

	"github.com/pedro-git-projects/go-chess/src/game"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	s := game.New().MarshalState()
	fmt.Fprintln(w, s)
}
