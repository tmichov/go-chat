package main

import (
	"fmt"
	"net/http"

	"github.com/tmichov/chat-go-react/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start connection")

		conn, err := websocket.Upgrade(w, r)
		if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
		}

		client := &websocket.Client{
				Conn: conn,
				Pool: pool,
		}

		pool.Register <- client
		client.Read()
}

func setupRoutes() {
		pool := websocket.NewPool()
		go pool.Start()

		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
				serveWs(pool, w, r)
		})
}

func main() {
		setupRoutes()
		http.ListenAndServe(":8000", nil)
}
