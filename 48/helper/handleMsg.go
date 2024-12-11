package helper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/index.html")
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	tracker := NewConversationTracker(10)
	for {
		msg := <-broadcast
		msg.Message = fmt.Sprintf("\n%s", msg.Message)
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}

		// AI bot response
		botResponse := generateBotResponse(tracker, msg.Message)
		if botResponse != "" {
			botMsg := Message{
				Username: "AI Bot",
				Message:  botResponse,
			}

			tracker.AddConversation(msg.Message, botResponse)
			for client := range clients {
				err := client.WriteJSON(botMsg)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func generateBotResponse(tracker *ConversationTracker, message string) string {
	message = strings.ToLower(message)
	if strings.Contains(message, "darmi") {
		listConverstion := tracker.GetHistory()
		return GetResponseAi(message, listConverstion)
	}
	return ""
}
