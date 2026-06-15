package main

import (
	"log"
	"net/http"

	"github.com/emanueldias01/mesh-backend/internal/handler"
)

func main() {
	http.HandleFunc("POST /rooms", handler.CreateRoom)
	http.HandleFunc("GET /rooms/{id}", handler.GetRoom)
	http.HandleFunc("GET /ws/rooms/{id}",handler.ConnectRoom)

	log.Println("Server running on 8080")
	http.ListenAndServe(":8080", nil)
}