package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var port = ":8080"

func main() {
	fmt.Printf("starting server on port %s\n", port)
	http.Handle("/create-room", websocket.Handler(receiveCreateRoom))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting server: %s\n", err)
	}
}
