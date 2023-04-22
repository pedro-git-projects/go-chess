package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var port = ":8080"

// I'm passing the create-room ws to the map
// I need the board ws in the map
func main() {
	srv := NewServer()
	http.Handle("/create-room", websocket.Handler(srv.receiveCreateRoom))
	http.Handle("/join-room", websocket.Handler(srv.receiveJoinRoom))
	http.Handle("/board", websocket.Handler(srv.receiveBoard))
	fmt.Printf("starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting server: %s\n", err)
	}
}
