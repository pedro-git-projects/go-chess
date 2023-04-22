package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/game"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
	"golang.org/x/net/websocket"
)

type GameServer struct {
	table         *game.Table
	clientsInRoom map[string]map[string]*websocket.Conn // roomID -> clientID -> ws
	mu            sync.RWMutex
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
		fmt.Printf("Client ID: %s\n", clientID)
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

func (s *GameServer) handleCalculateLegalMoves(ws *websocket.Conn, r *BoardRequest) {
	roomID := r.RoomID
	if r.Coordinate != nil && r.Message == "calc" && s.table.HasKey(roomID) {
		c, err := utils.CoordFromStr(*r.Coordinate)
		if err != nil {
			log.Println("Error converting object to coordinate:", err)
			return
		}
		gameState := s.table.Game(roomID)
		if gameState == nil {
			res := CalculateResponse{
				LegalMovements: "",
				Error:          fmt.Sprintf("Invalid game state"),
			}
			err := websocket.JSON.Send(ws, res)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
			}
			return
		}
		color := gameState.PieceColor(c)
		l := gameState.LegalMovements(c, color)
		marshaled, err := json.Marshal(l)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshalling response: %s\n", err)
		}
		res := CalculateResponse{
			LegalMovements: string(marshaled),
			Error:          "",
		}
		//	err = s.broadcastMessageInRoom(roomID, res) // Error origin
		err = websocket.JSON.Send(ws, res)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending response: %s\n", err)
			return
		}
	}
}

func (s *GameServer) handleRender(ws *websocket.Conn, r *BoardRequest) {
	roomID := r.RoomID
	if r.Message != "render" || !s.table.HasKey(roomID) {
		res := RenderBoardResponse{
			GameState: "",
			Error:     fmt.Sprintf("Could not render board, invalid data"),
		}
		err := websocket.JSON.Send(ws, res)
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
		err := s.broadcastMessageInRoom(roomID, res)
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
	err := websocket.JSON.Send(ws, resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "render::error 2 receiving message: %s\n", err)
	}
}

func (s *GameServer) receiveBoard(ws *websocket.Conn) {
	for {
		r := new(BoardRequest)
		err := websocket.JSON.Receive(ws, r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "board::error 1 receiving message: %s\n", err)
			return
		}
		switch r.Message {
		case "render":
			s.handleRender(ws, r)
		case "calc":
			s.handleCalculateLegalMoves(ws, r)
		}
	}
}

// broadcastMessage sends a message to all clients in a room
// is the error
func (s *GameServer) broadcastMessageInRoom(roomID string, msg interface{}) error {
	s.mu.RLock()
	clients := s.clientsInRoom[roomID]
	fmt.Printf("Clients in room: %v\n", clients)
	s.mu.RUnlock()

	for clientID, ws := range clients {
		fmt.Printf("Client ID in broadcast: %s\n", clientID)
		err := websocket.JSON.Send(ws, msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message to client %s: %s\n", clientID, err)
			// remove the client from the list of clients in the room
			s.mu.Lock()
			delete(clients, clientID)
			s.mu.Unlock()
		}
	}

	return nil
}

// addClientToRoom checks if a the key that has roomID as a key exists in the clientsInRoom map
// if it doesn't a new empty map will be created and assigned to that roomID
// then the clientID is added to the map with the connection as its value and roomID
// as a key to its map.
func (s *GameServer) addClientToRoom(roomID, clientID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.clientsInRoom[roomID]; !ok {
		s.clientsInRoom[roomID] = make(map[string]*websocket.Conn)
	}
	s.clientsInRoom[roomID][clientID] = conn
}

// setClientConnection updates the connection in the map
func (s *GameServer) setClientConnection(roomID, clientID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.clientsInRoom[roomID]; ok {
		s.clientsInRoom[roomID][clientID] = conn
	}
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
