package handler

import (
	"encoding/json"
	"net/http"

	"github.com/emanueldias01/mesh-backend/internal/room"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	room := room.CreateRoom()

	response := map[string]string{
		"roomId": room.ID,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {

	roomID := r.PathValue("id")

	room, ok := room.GetRoom(roomID)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(room)
}