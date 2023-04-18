package main

import (
	"fmt"
	"os"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/game"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
	"golang.org/x/net/websocket"
)

type GameServer struct {
	table *game.Table
}

func NewServer() *GameServer {
	return &GameServer{
		table: game.NewTable(),
	}
}

func (s *GameServer) receiveCreateRoom(ws *websocket.Conn) {
	msg := new(CreateRoomRequest)
	err := websocket.JSON.Receive(ws, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error receiving message: %s\n", err)
	}
	if msg.Message == "create" {
		roomID := utils.GenerateRoomId()
		clientID := utils.GenerateRoomId()
		gameState := game.New()
		s.table.SetGame(roomID, gameState)
		err := s.table.Game(roomID).AddClient(game.NewClient(clientID))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
		}
		resp := CreateRoomResponse{
			RoomID:   roomID,
			ClientID: clientID,
		}
		err = websocket.JSON.Send(ws, resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
		}
	}
}

// receiveJoinRoom expects to receive a JSON object containing
// the message join and a roomID
// if the roomID is in the game map and there is an empty slot
// then the client joins that room
func (s *GameServer) receiveJoinRoom(ws *websocket.Conn) {
	r := new(JoinRoomRequest)
	err := websocket.JSON.Receive(ws, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error receiving message: %s\n", err)
	}
	roomID := r.RoomID
	if r.Message == "join" && s.table.HasKey(roomID) {
		clientID := utils.GenerateRoomId()
		err := s.table.Game(roomID).AddClient(game.NewClient(clientID))
		gameState := s.table.Game(roomID)
		if err == nil {
			s.table.SetGame(roomID, gameState)
			resp := JoinRoomResponse{
				RoomID:   roomID,
				ClientID: clientID,
			}
			err = websocket.JSON.Send(ws, resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
		}
	}
}
