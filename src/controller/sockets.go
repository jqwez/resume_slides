package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type SocketHub struct {
	clients    map[*websocket.Conn]bool
	clientsMux sync.Mutex
	state      State
	upgrader   websocket.Upgrader
}

func NewSocketHub() *SocketHub {
	return &SocketHub{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *SocketHub) RegisterRoutes() {
	http.HandleFunc("/ws", s.HandleSocket)
}

func (s *SocketHub) HandleSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	s.clientsMux.Lock()
	s.clients[conn] = true
	s.clientsMux.Unlock()
	s.sendState(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.removeClient(conn)
			break
		}
		s.ProcessMessage(msg)
	}
}

func (s *SocketHub) sendState(client *websocket.Conn) {
	s.clientsMux.Lock()
	defer s.clientsMux.Unlock()

	stateJSON := []byte(fmt.Sprintf(`{"state": {"slideshow": %s, "slide": %s}}`, fmt.Sprint(s.state.Slideshow), fmt.Sprint(s.state.Slide)))
	err := client.WriteMessage(websocket.TextMessage, stateJSON)
	if err != nil {
		log.Println("Error sending state to client:", err)
	}
}

func (s *SocketHub) broadcastState() {
	s.clientsMux.Lock()
	clients := make(map[*websocket.Conn]bool, len(s.clients))
	for client := range s.clients {
		clients[client] = true
	}
	s.clientsMux.Unlock()

	for client := range s.clients {
		s.sendState(client)
	}
}

func (s *SocketHub) removeClient(client *websocket.Conn) {
	s.clientsMux.Lock()
	defer s.clientsMux.Unlock()

	delete(s.clients, client)
}

func (s *SocketHub) setState(newState State) {
	s.state = newState
	s.broadcastState()
}

func (s *SocketHub) ProcessMessage(msg []byte) {
	fmt.Println("received message: ", string(msg))
	state := State{}
	err := json.Unmarshal(msg, &state)
	if err != nil {
		log.Println("Failed to parse State")
	}
	s.setState(state)
}
