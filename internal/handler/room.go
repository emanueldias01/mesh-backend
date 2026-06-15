package handler

import (
	"encoding/json"
	"net/http"

	"github.com/emanueldias01/mesh-backend/internal/room"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	room := room.CreateRoom()

	response := map[string]string {
		"roomId" : room.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(response)
}