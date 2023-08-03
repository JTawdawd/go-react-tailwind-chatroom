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
	var req LoginRequest
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

// func getUser1() {
// 	var (
// 		id       int
// 		username string
// 	)
// 	rows, err := DB.Query("select id, username from user_account where id = $1", 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(&id, &username)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(id, username)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
