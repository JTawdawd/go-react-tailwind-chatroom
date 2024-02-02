package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connections = make([]*websocket.Conn, 0)
	connMu      sync.Mutex
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if handlerDefinitions[r.URL.Path] == nil {
		http.Error(w, "handler not defined", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response, err := handlerDefinitions[r.URL.Path](body)
	log.Print(err)
	if err != nil {
		response, _ = json.Marshal(map[string]interface{}{
			"status": "Error",
			"Error":  "handler failed: " + err.Error(),
		})
	}

	var res Response
	json.Unmarshal(response, &res)
	if res.Message == "Created Message" {
		for _, connection := range connections {
			connection.WriteMessage(1, []byte("New message"))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if handlerDefinitions[r.URL.Path] == nil {
		http.Error(w, "handler not defined", http.StatusBadRequest)
		return
	}
	response, err := handlerDefinitions[r.URL.Path]([]byte{})
	if err != nil {
		response, _ = json.Marshal(map[string]interface{}{
			"status": "Error",
			"Error":  "handler failed: " + err.Error(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved websock connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connMu.Lock()
	connections = append(connections, conn)
	connMu.Unlock()

	defer func() {
		connMu.Lock()
		defer connMu.Unlock()

		for i, c := range connections {
			if c == conn {
				connections = append(connections[:i], connections[i+1:]...)
				break
			}
		}
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			break
		}
	}
}
