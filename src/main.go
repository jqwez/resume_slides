package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"

	"main/routes"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Go")
	routes.RegisterRoutes()
	port := ":8000"
	http.ListenAndServe(port, nil)
}