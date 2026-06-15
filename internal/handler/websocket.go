package handler

import (
	"net/http"

	"github.com/emanueldias01/mesh-backend/internal/room"
	"github.com/emanueldias01/mesh-backend/internal/websocket"
	gorilla "github.com/gorilla/websocket"
)

var upgrader = gorilla.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Broadcast(currentRoom *room.Room, payload []byte) {
	currentRoom.Mutex.RLock()
	defer currentRoom.Mutex.RUnlock()

	for _, client := range currentRoom.Clients {
		if err := client.Conn.WriteMessage(gorilla.TextMessage, payload); err != nil {
			continue
		}
	}
}

func Listen(currentRoom *room.Room, client *websocket.Client) {

	defer func() {

		currentRoom.Mutex.Lock()
		delete(currentRoom.Clients, client.ID)
		currentRoom.Mutex.Unlock()

		Broadcast(
			currentRoom,
			[]byte(client.ID+" saiu da sala"),
		)

		client.Conn.Close()
	}()

	for {
		_, payload, err := client.Conn.ReadMessage()

		if err != nil {
			break
		}

		Broadcast(currentRoom, payload)
	}
}

func ConnectRoom(w http.ResponseWriter, r *http.Request) {

	roomID := r.PathValue("id")
	userID := r.URL.Query().Get("userId")

	if userID == "" {
		http.Error(w, "userId required", http.StatusBadRequest)
		return
	}

	currentRoom, ok := room.GetRoom(roomID)

	if !ok {
		http.NotFound(w, r)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	client := &websocket.Client{
		ID:   userID,
		Conn: conn,
	}

	currentRoom.Mutex.Lock()
	currentRoom.Clients[userID] = client
	currentRoom.Mutex.Unlock()

	Broadcast(
		currentRoom,
		[]byte(userID+" entrou na sala"),
	)

	Listen(currentRoom, client)
}