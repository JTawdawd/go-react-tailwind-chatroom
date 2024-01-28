package handler

import (
	"encoding/json"
	"errors"
	"log"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(data []byte) ([]byte, error) {
	var decodedData registerRequest
	err := json.Unmarshal(data, &decodedData)
	if err != nil {
		return nil, err
	}

	log.Printf("Received create request:\nUsername: %s\nPassword: %s\n", decodedData.Username, decodedData.Password)

	err = insert("CreateUser", decodedData.Username, decodedData.Password)
	if err != nil {
		return []byte("Failed to create user"), err
	}

	rows, err := query("GetUser", decodedData.Username, decodedData.Password)
	if err != nil {
		return nil, err
	}
	var (
		id       int
		username string
	)
	for rows.Next() {
		err = rows.Scan(&id, &username)
		if err != nil {
			return nil, err
		}
	}
	if id == 0 && username == "" {
		return nil, errors.New("no user found")
	}
	return json.Marshal(map[string]interface{}{
		"status":   "Created user",
		"id":       id,
		"username": username,
	})
}
