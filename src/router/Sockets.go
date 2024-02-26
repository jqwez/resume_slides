package router

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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
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

func processMessage(message string) SocketResponse {
	log.Println(message)
	return SocketResponse{
		Action:  "Action",
		Message: message,
	}
}
