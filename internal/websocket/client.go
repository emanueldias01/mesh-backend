package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ID string `json:"id"`
	Conn *websocket.Conn
}