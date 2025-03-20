package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan SensorData)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	db *sql.DB
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket Read error:", err)
			delete(clients, conn)
			return
		}
	}
}

func handleWebSocketBroadcast() {
	for {
		data := <-broadcast
		for client := range clients {
			jsonData, _ := json.Marshal(data)
			client.WriteMessage(websocket.TextMessage, jsonData)
		}
	}
}
