package main

type CreateRoomRequest struct {
	Message string `json:"message"`
}

type CreateRoomResponse struct {
	RoomID   string `json:"room_id"`
	ClientID string `json:"client_id"`
}

type JoinRoomRequest struct {
	Message string `json:"message"`
	RoomID  string `json:"room_id"`
}

type JoinRoomResponse struct {
	RoomID   string `json:"room_id"`
	ClientID string `json:"client_id"`
	Error    string `json:"error"`
}

type RenderBoardRequest struct {
	Message string `json:"message"`
	RoomID  string `json:"room_id"`
}

type RenderBoardResponse struct {
	GameState string `json:"state"`
	Error     string `json:"error"`
}

type CalculateRequest struct {
	Message    string `json:"message"`
	Coodrinate string `json:"coordinate"`
	RoomID     string `json:"room_id"`
}

type CalculateResponse struct {
	LegalMovements string `json:"legal_movements"`
	Error          string `json:"error"`
}
