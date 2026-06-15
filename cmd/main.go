package main

import (
	"log"
	"net/http"

	"github.com/emanueldias01/mesh-backend/internal/handler"
	"github.com/emanueldias01/mesh-backend/internal/websocket"
)

func main() {
	http.HandleFunc("/rooms", handler.CreateRoom)
	http.HandleFunc("/ws", websocket.HandleWS)
	log.Println("Server running on 8080")
	http.ListenAndServe(":8080", nil)
}