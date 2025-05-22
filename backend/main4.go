// File: backend/websocket_server.go
package main

import (
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Izinkan CORS
}

// Hub untuk mengelola koneksi WebSocket
type Hub struct {
	clients map[*websocket.Conn]bool
}

var hub = Hub{
	clients: make(map[*websocket.Conn]bool),
}

func main() {
	// Inisialisasi MQTT Client
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:8888")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT connection error:", token.Error())
	}

	// Subscribe ke topik Arduino
	client.Subscribe("arduino/data", 0, func(c mqtt.Client, m mqtt.Message) {
		// Broadcast ke semua WebSocket clients
		for conn := range hub.clients {
			err := conn.WriteMessage(websocket.TextMessage, m.Payload())
			if err != nil {
				log.Println("WebSocket write error:", err)
				conn.Close()
				delete(hub.clients, conn)
			}
		}
	})

	// Setup WebSocket endpoint
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// Tambahkan client ke hub
	hub.clients[conn] = true

	// Keep connection alive
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			delete(hub.clients, conn)
			break
		}
	}
}
