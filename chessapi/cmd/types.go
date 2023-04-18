package main

type CreateRoomRequest struct {
	Message string `json:"message"`
}

type CreateRoomResponse struct {
	RoomID   string `json:"room_id"`
	ClientID string `json:"client_id"`
}
