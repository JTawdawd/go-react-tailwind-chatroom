package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
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
	if strings.HasPrefix(res.Message, "Created Message:") {
		chatroomManager.SendMessageToChatroom(strings.Replace(res.Message, "Created Message:", "", -1), "New message")
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
	chatroomID := r.URL.Query().Get("chatroomID")
	if chatroomID == "" {
		log.Println("Missing chatroomID in query parameters")
		http.Error(w, "Missing chatroomID", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	chatroomManager.AddConnectionToChatroom(string(chatroomID), conn)
	defer chatroomManager.RemoveConnectionFromChatroom(string(chatroomID), conn)

	for {
		messageType, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		if messageType == websocket.CloseMessage {
			break
		}
	}
}
