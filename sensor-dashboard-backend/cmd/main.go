package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	sensordashboard "github.com/manan1979/sensor-dashboard"
)

func main() {

	cfg := sensordashboard.LoadConfiguration()

	var err error
	db, err = sql.Open("mysql", cfg.DB.CreateDSN())
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", cfg.Bind.HTTP)
	if err != nil {
		log.Fatal("TCP server error:", err)
	}
	defer listener.Close()

	log.Println("TCP Server started on port 9000")

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("TCP Connection error:", err)
				continue
			}
			go handleTCPConnection(conn)
		}
	}()

	go handleWebSocketBroadcast()

	http.HandleFunc("/ws", handleWebSocket)
	log.Println("WebSocket server started on ws://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
