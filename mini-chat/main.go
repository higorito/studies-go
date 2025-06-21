package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	UserName string `json:"username"`
	Message  string `json:"message"`
}

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	// websocket.DefaultDialer.HandshakeTimeout: 10 * 60 * 1000, // 10 minutes
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnection)

	go handleMessage()

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Could not start server: " + err.Error())
	}
}

var handleConnection = func(w http.ResponseWriter, r *http.Request) {
	webs, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
		return
	}

	defer webs.Close()
	clients[webs] = true

	for {
		var msg Message
		err := webs.ReadJSON(&msg)
		if err != nil {
			delete(clients, webs)
			break
		}

		broadcast <- msg
	}
}

var handleMessage = func() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				delete(clients, client)
				client.Close()
			}
		}
	}
}
