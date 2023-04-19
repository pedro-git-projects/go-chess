package game

import "github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/piece"

type Client struct {
	id    string
	color piece.Color
}

func NewClient(clientID string) *Client {
	return &Client{
		id:    clientID,
		color: piece.None,
	}
}

func (c *Client) Color() piece.Color {
	return c.color
}
