package main

import (
	"chatroom"
	"net/http"
	"strings"
)

var chatroomManager *chatroom.ChatroomManager

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "static") {
		http.ServeFile(w, r, "../frontend/build"+r.URL.Path)
		return
	}

	if !strings.HasPrefix(r.URL.Path, "/api") {
		http.ServeFile(w, r, "../frontend/build/index.html")
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func serveReactApp(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("../frontend/build/index.html"))
	http.StripPrefix("/", fs).ServeHTTP(w, r)
}

func main() {
	chatroomManager = chatroom.CreateChatroom()

	defineHandlers()

	http.HandleFunc("/websocket/connect", handleWebsocket)
	http.Handle("/", &myHandler{})
	http.ListenAndServe(":8080", nil)
}
