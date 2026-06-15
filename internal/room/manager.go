package room

import (
	"sync"

	"github.com/emanueldias01/mesh-backend/internal/websocket"
)

var (
	rooms = make(map[string]*Room)
	mutex sync.RWMutex
)

func CreateRoom() *Room{
	mutex.Lock()
	defer mutex.Unlock()

	code := GenerateCode()

	room := &Room{
		ID: code,
		Clients: make(map[string]*websocket.Client),
	}

	rooms[code] = room

	return room
}

func GetRoom(code string) (*Room, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	room, ok := rooms[code]

	return room, ok
}