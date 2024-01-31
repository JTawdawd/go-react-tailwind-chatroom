package handler

import (
	"encoding/json"
)

type Message struct {
	Username  string `json:"username"`
	UserId    string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
}

type chatroomRequest struct {
	ID int `json:"id"`
}

func GetChatroom(data []byte) ([]byte, error) {

	var decodedData chatroomRequest

	err := json.Unmarshal(data, &decodedData)
	if err != nil {
		return nil, err
	}
	rows, err := query("GetMessages", decodedData.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var message Message
		err := rows.Scan(&message.Username, &message.UserId, &message.Content, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	response, err := json.Marshal(map[string]interface{}{
		"status":   "Success",
		"messages": messages,
	})
	if err != nil {
		return nil, err
	}
	return response, nil

	// open websocket?
}
