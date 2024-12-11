package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Order represents an order update message
type Order struct {
	ID          int       `json:"id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type OrderPayload struct {
	ID          int       `json:"id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

var data []string

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Handle WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer ws.Close()
	clients[ws] = true

	for {
		_, msgBytes, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, ws)
			break
		}

		var msg Message
		err = json.Unmarshal(msgBytes, &msg)
		if err != nil {
			log.Printf("Error parsing message: %v", err)
			continue
		}

		// Handle chat messages
		if msg.Type == "chat" {
			responseMsg := handleChat(msg)
			broadcast <- responseMsg
		}
	}
}

// Handle messages for broadcasting
func handleMessages() {
	for {
		msg := <-broadcast
		// Pastikan pesan order diproses dengan benar
		message, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// Generate a chatbot response
func handleChat(msg Message) Message {
	// Ekstrak payload dari interface{}
	payload, ok := msg.Payload.(map[string]interface{})
	if !ok {
		log.Printf("Invalid payload format for chat message")
		return Message{
			Type: "chat",
			Payload: map[string]interface{}{
				"question": "",
				"response": "Error: Invalid message format",
			},
		}
	}

	// Ambil pertanyaan dari payload
	question, ok := payload["question"].(string)
	if !ok {
		question = "No question provided"
	}

	// Simple chatbot logic (replace this with actual AI/bot logic)
	response := GetResponseAi(question, data)

	// Kembalikan pesan dengan struktur yang sesuai
	return Message{
		Type: "chat",
		Payload: map[string]interface{}{
			"question": question,
			"response": response,
		},
	}
}

// Simulate order updates
func simulateOrderUpdates() {
	orderID := 1
	statuses := []string{"pending", "ongoing", "canceled", "completed"}
	for {
		statusChoice := statuses[rand.Intn(len(statuses))]
		randomAmount := float64(50 + rand.Intn(495))

		orderPayload := OrderPayload{
			ID:          orderID,
			TotalAmount: randomAmount,
			Status:      statusChoice,
			CreatedAt:   time.Now(),
		}

		msg := Message{
			Type:    "order",
			Payload: orderPayload,
		}

		broadcast <- msg
		orderID++
		time.Sleep(3 * time.Second)
		log.Printf("Successfully created Order #%d with status '%s' and amount %.2f", orderID, statusChoice, randomAmount)
		data = append(data, fmt.Sprintf("Successfully created Order #%d with status '%s' and amount %.2f", orderID, statusChoice, randomAmount))
	}
}

// Serve the homepage
func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	// Route handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ws", handleConnections)

	// Goroutines for handling messages and simulating updates
	go handleMessages()
	go simulateOrderUpdates()

	// Start the server
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
