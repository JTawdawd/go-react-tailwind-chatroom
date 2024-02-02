package handler

import (
	"encoding/json"
	"errors"
)

type MessageRequest struct {
	ChatroomId string `json:"chatroomid"`
	UserId     string `json:"createdby"`
	Content    string `json:"content"`
	CreatedAt  string `json:"createdat"`
}

func CreateMessage(data []byte) ([]byte, error) {
	var decodedData MessageRequest
	err := json.Unmarshal(data, &decodedData)
	if err != nil {
		return nil, errors.New("Failed to decode data")
	}

	err = insert("CreateMessage", decodedData.ChatroomId, decodedData.UserId, decodedData.CreatedAt, decodedData.Content)
	if err != nil {
		return nil, errors.New("Failed to create message")
	}

	return json.Marshal(map[string]interface{}{
		"status":  "Success",
		"message": "Created Message",
	})
}
