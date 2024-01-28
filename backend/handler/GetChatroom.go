package handler

import (
	"encoding/json"
)

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	CreatedBy string `json:"createdby"`
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
		err := rows.Scan(&message.ID, &message.Content, &message.CreatedBy, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return json.Marshal(messages)

	// open websocket?
}
