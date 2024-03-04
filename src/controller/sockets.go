package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type SocketResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

func NewSocketResponse(action string, message string) *SocketResponse {
	return &SocketResponse{
		Action:  action,
		Message: message,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		response := processMessage(string(message))
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteMessage(messageType, jsonResponse)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func processMessage(message string) *SocketResponse {
	log.Println("Socket Message: ", message)
	return NewSocketResponse("Action", message)
}
