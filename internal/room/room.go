package room

import (
	"sync"

	"github.com/emanueldias01/mesh-backend/internal/websocket"
)

type Room struct {
	ID string `json:"id"`
	Clients map[string]*websocket.Client `json:"-"`
	Mutex   sync.RWMutex          `json:"-"`
}

type RoomRequest struct {
	ID string `json:"id"`
}