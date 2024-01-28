package handler

import (
	"encoding/json"
	"fmt"
	"log"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(data []byte) ([]byte, error) {
	var decodedData LoginRequest
	if err := json.Unmarshal(data, &decodedData); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	log.Printf("Received login request:\nUsername: %s\nPassword: %s\n", decodedData.Username, decodedData.Password)

	rows, err := query("GetUser", decodedData.Username, decodedData.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}
	defer rows.Close()

	var (
		id       int
		username string
	)

	if rows.Next() {
		if err := rows.Scan(&id, &username); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}

	log.Printf("Login successful:\nUsername: %s\nID: %d\n", username, id)

	response := map[string]interface{}{
		"status":   "Login request received",
		"id":       id,
		"username": username,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %w", err)
	}

	return jsonData, nil
}
