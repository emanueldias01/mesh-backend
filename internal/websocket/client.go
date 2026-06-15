package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string `json:"id"`
	Conn *websocket.Conn
	Mu   sync.Mutex
}

func (c *Client) WriteJSON(v interface{}) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Conn.WriteJSON(v)
}