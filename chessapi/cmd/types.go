package main

type CreateRoomRequest struct {
	Message string `json:"message"`
}

type CreateRoomResponse struct {
	ClientID string `json:"client_id"`
}
