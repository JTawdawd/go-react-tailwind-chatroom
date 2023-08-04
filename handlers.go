package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type chatroomRequest struct {
	ID int `json:"id"`
}

type Chatroom struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	CreatedBy string `json:"createdby"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Received login request:\nUsername: %s\nPassword: %s\n", req.Username, req.Password)
	row := DB.QueryRow("select id, username from user_account where username = $1 AND (password is not null and password = crypt($2, password))", req.Username, req.Password)
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"error\": \"No user found\"}"))
		return
	}

	var (
		id       int
		username string
	)
	err = row.Scan(&id, &username)
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"error\": \"No user found\"}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"status\": \"Login request received\", \"id\": \"%d\", \"username\": \"%s\"}", id, username)))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Received create request:\nUsername: %s\nPassword: %s\n", req.Username, req.Password)

	stmt, err := DB.Prepare("INSERT INTO user_account(username, password) VALUES($1, crypt($2, gen_salt('bf')))")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(req.Username, req.Password)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login request received"))
}

func GetChatroom(w http.ResponseWriter, r *http.Request) {
	var req chatroomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// return an array of message objects with createdBy, createdAt, content
	rows, err := DB.Query("SELECT id, content, createdby, createdat FROM message WHERE chatroomid = $1", req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var message Message
		err := rows.Scan(&message.ID, &message.Content, &message.CreatedBy, &message.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}

	// Convert the chatrooms slice to JSON format.
	chatRoomsJSON, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(chatRoomsJSON)

	// open websocket?
}

func Chatrooms(w http.ResponseWriter, r *http.Request) {
	// Query the database to fetch all chatrooms.
	rows, err := DB.Query("SELECT id, title FROM chatroom")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var chatrooms []Chatroom

	for rows.Next() {
		var chatroom Chatroom
		err := rows.Scan(&chatroom.ID, &chatroom.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		chatrooms = append(chatrooms, chatroom)
	}

	// Convert the chatrooms slice to JSON format.
	chatroomsJSON, err := json.Marshal(chatrooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(chatroomsJSON)
}
