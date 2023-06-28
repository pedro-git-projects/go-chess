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
	connStr := os.Getenv("DB_CONNECTION_STRING")

	if connStr == "" {
		fmt.Println("DB_CONNECTION_STRING environment variable not set")
		return
	}

	db, err := auth.NewDatabase(connStr)
	if err != nil {
		log.Fatal(err)
	}
	authService, err := auth.NewAuthService(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := NewServer()
	http.HandleFunc("/login", auth.CorsMiddleware(auth.HandleLogin(authService)))
	http.HandleFunc("/register", auth.CorsMiddleware(auth.HandleRegistration(authService)))
	http.HandleFunc("/change-password", auth.CorsMiddleware(auth.HandleChangePassword(authService)))
	http.HandleFunc("/delete", auth.CorsMiddleware(auth.HandleDeleteUser(authService)))
	http.HandleFunc("/signout", auth.CorsMiddleware(auth.HandleSignout(authService)))
	http.Handle("/game", websocket.Handler(srv.gameLoop))
	fmt.Printf("starting server on port %s\n", port)
	//	err = http.ListenAndServeTLS(port, "ssl/certs/cert.pem", "ssl/certs/key.pem", nil)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting server: %s\n", err)
	}
}
