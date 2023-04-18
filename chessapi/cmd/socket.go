package main

import (
	"fmt"
	"os"

	"golang.org/x/net/websocket"
)

// receiveCreateRoom expects to receive a
// JSON object {"message": "create"}
// if such an object is received,
// a new game Room is created
func receiveCreateRoom(ws *websocket.Conn) {
	msg := new(CreateRoomMessage)
	err := websocket.JSON.Receive(ws, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error receiving message: %s\n", err)
	}
	fmt.Println(msg.Message)
}
