package handler

import (
	"encoding/json"
	"errors"
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

	rows, err := query("CheckUserExist", decodedData.Username)
	if err != nil {
		return nil, errors.New("Failed to check for current users")
	}
	if rows.Next() {
		return nil, errors.New("User already exists")
	}

	err = insert("CreateUser", decodedData.Username, decodedData.Password)
	if err != nil {
		return []byte("Failed to create user"), err
	}

	rows, err = query("GetUser", decodedData.Username, decodedData.Password)
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
		"status":   "Success",
		"message":  "Created User",
		"id":       id,
		"username": username,
	})
}
