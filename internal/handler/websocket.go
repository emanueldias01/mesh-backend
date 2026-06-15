package handler

import (
	"encoding/json"
	"log"
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

type SignalMessage struct {
	Type    string          `json:"type"`
	From    string          `json:"from,omitempty"`
	To      string          `json:"to,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

func sendTo(currentRoom *room.Room, targetID string, msg SignalMessage) {
	currentRoom.Mutex.RLock()
	client, ok := currentRoom.Clients[targetID]
	currentRoom.Mutex.RUnlock()

	if !ok {
		return
	}
	if err := client.WriteJSON(msg); err != nil {
		log.Println("error to send to", targetID, ":", err)
	}
}

func broadcastExcept(currentRoom *room.Room, exceptID string, msg SignalMessage) {
	currentRoom.Mutex.RLock()
	defer currentRoom.Mutex.RUnlock()

	for id, client := range currentRoom.Clients {
		if id == exceptID {
			continue
		}
		if err := client.WriteJSON(msg); err != nil {
			log.Println("Error in broadcast to", id, ":", err)
		}
	}
}

func Listen(currentRoom *room.Room, client *websocket.Client) {
	defer func() {
		currentRoom.Mutex.Lock()
		delete(currentRoom.Clients, client.ID)
		currentRoom.Mutex.Unlock()

		broadcastExcept(currentRoom, client.ID, SignalMessage{
			Type: "user-left",
			From: client.ID,
		})

		client.Conn.Close()
	}()

	for {
		var msg SignalMessage
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		msg.From = client.ID

		switch msg.Type {
		case "chat":
			currentRoom.Mutex.RLock()
			for _, c := range currentRoom.Clients {
				c.WriteJSON(msg)
			}
			currentRoom.Mutex.RUnlock()

		case "offer", "answer", "ice-candidate":
			if msg.To != "" {
				sendTo(currentRoom, msg.To, msg)
			}

		default:
		}
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
	existingIDs := make([]string, 0, len(currentRoom.Clients))
	for id := range currentRoom.Clients {
		existingIDs = append(existingIDs, id)
	}
	currentRoom.Clients[userID] = client
	currentRoom.Mutex.Unlock()

	existingPayload, _ := json.Marshal(existingIDs)
	client.WriteJSON(SignalMessage{
		Type:    "existing-users",
		Payload: existingPayload,
	})

	broadcastExcept(currentRoom, userID, SignalMessage{
		Type: "user-joined",
		From: userID,
	})

	Listen(currentRoom, client)
}