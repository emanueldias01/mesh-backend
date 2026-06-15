package main

import (
	"net/http"

	"github.com/emanueldias01/mesh-backend/internal/handler"
)

func main() {
	http.HandleFunc("/rooms", handler.CreateRoom)

	http.ListenAndServe(":8080", nil)
}