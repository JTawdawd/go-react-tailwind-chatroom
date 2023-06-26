package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login request received"))
}
