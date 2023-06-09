package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/auth"
	"golang.org/x/net/websocket"
)

var port = ":8080"

func main() {
	db, err := auth.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	authService, err := auth.NewAuthService(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := NewServer()
	http.HandleFunc("/login", auth.HandleLogin(authService))
	http.HandleFunc("/register", auth.HandleRegistration(authService))
	http.Handle("/game", websocket.Handler(srv.gameLoop))
	fmt.Printf("starting server on port %s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting server: %s\n", err)
	}
}
