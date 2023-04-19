package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/game"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
	"golang.org/x/net/websocket"
)

type GameServer struct {
	table         *game.Table
	clientsInRoom map[string]map[string]*websocket.Conn // roomID -> clientID -> ws
	mu            sync.Mutex
}

func NewServer() *GameServer {
	return &GameServer{
		table:         game.NewTable(),
		clientsInRoom: make(map[string]map[string]*websocket.Conn),
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
		gameState := game.NewGame()
		s.table.SetGame(roomID, gameState)
		// try to add client to game
		err := s.table.Game(roomID).AddClient(game.NewClient(clientID))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
		}
		// add client to room
		s.addClientToRoom(roomID, clientID, ws)
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
	// receive request
	r := new(JoinRoomRequest)
	err := websocket.JSON.Receive(ws, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error receiving message: %s\n", err)
	}
	roomID := r.RoomID
	if r.Message == "join" && s.table.HasKey(roomID) {
		clientID := utils.GenerateRoomId()
		gameState := s.table.Game(roomID)
		// check if gameState has been correctly populated
		if gameState == nil {
			resp := JoinRoomResponse{
				RoomID:   "",
				ClientID: "",
				Error:    fmt.Sprintf("invalid room ID: %s", roomID),
			}
			err = websocket.JSON.Send(ws, resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
			}
			return
		}
		// try to add client to game
		err := gameState.AddClient(game.NewClient(clientID))
		if err != nil {
			resp := JoinRoomResponse{
				RoomID:   "",
				ClientID: "",
				Error:    err.Error(),
			}
			err = websocket.JSON.Send(ws, resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
			}
			return
		}
		// add client to room
		s.addClientToRoom(roomID, clientID, ws)
		s.table.SetGame(roomID, gameState)
		resp := JoinRoomResponse{
			RoomID:   roomID,
			ClientID: clientID,
			Error:    "",
		}
		err = websocket.JSON.Send(ws, resp)
		if err != nil {
			resp := JoinRoomResponse{
				RoomID:   "",
				ClientID: "",
				Error:    err.Error(),
			}
			err = websocket.JSON.Send(ws, resp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
			}
		}
	} else {
		resp := JoinRoomResponse{
			RoomID:   "",
			ClientID: "",
			Error:    fmt.Sprintf("invalid room ID: %s", roomID),
		}
		err = websocket.JSON.Send(ws, resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
		}
	}
}

// receiveRenderBoard expects a {"message":"render", "room_id":"id"}
// JSON object, if it is recived
// it checks if the roomID is on the game table
// and if it is, it responds with {"state":"[...]", and "error":"error"}
func (s *GameServer) receiveRenderBoard(ws *websocket.Conn) {
	r := new(RenderBoardRequest)
	err := websocket.JSON.Receive(ws, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "render::error 1 receiving message: %s\n", err)
		return
	}
	roomID := r.RoomID
	if r.Message != "render" || !s.table.HasKey(roomID) {
		res := RenderBoardResponse{
			GameState: "",
			Error:     fmt.Sprintf("Could not render board, invalid data"),
		}
		err = websocket.JSON.Send(ws, res)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to send response: %s\n", err)
		}
		return
	}
	gameState := s.table.Game(roomID)
	if gameState == nil {
		res := RenderBoardResponse{
			GameState: "",
			Error:     fmt.Sprintf("Could not render board: invalid game state"),
		}
		err := s.broadcastMessage(res, roomID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error broadcasting message: %s\n", err)
		}
		return
	}
	m := gameState.MarshalState()
	resp := RenderBoardResponse{
		GameState: m,
		Error:     "",
	}
	err = websocket.JSON.Send(ws, resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "render::error 2 receiving message: %s\n", err)
	}
}

// broadcastMessage sends a message to all clients in a room
func (s *GameServer) broadcastMessage(msg interface{}, roomID string) error {
	if clients, ok := s.clientsInRoom[roomID]; ok {
		for _, conn := range clients {
			if err := websocket.JSON.Send(conn, msg); err != nil {
				return fmt.Errorf("error sending message: %s", err)
			}
		}
		return nil
	}
	return fmt.Errorf("invalid room ID: %s", roomID)
}

// addClientToRoom is a concurrent safe method for adding clients to a room
func (s *GameServer) addClientToRoom(roomID, clientID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.clientsInRoom[roomID]; !ok {
		s.clientsInRoom[roomID] = make(map[string]*websocket.Conn)
	}
	s.clientsInRoom[roomID][clientID] = conn
}

// removeClientFromRoom is a concurrent safe method for removing clients from a room
func (s *GameServer) removeClientFromRoom(roomID, clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.clientsInRoom[roomID]; ok {
		delete(s.clientsInRoom[roomID], clientID)
		if len(s.clientsInRoom[roomID]) == 0 {
			delete(s.clientsInRoom, roomID)
		}
	}
}
