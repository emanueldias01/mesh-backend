package room

import (
	"sync"

	"github.com/emanueldias01/mesh-backend/internal/websocket"
)

type Room struct {
	ID string `json:"id"`
	Clients map[string]*websocket.Client
	Mutex   sync.RWMutex
}

type RoomRequest struct {
	ID string `json:"id"`
}