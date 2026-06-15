package room

import "sync"

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