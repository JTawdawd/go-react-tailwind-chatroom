package handler

import (
	"encoding/json"
	"errors"
)

type Chatroom struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetChatrooms(data []byte) ([]byte, error) {

	// var decodedData Chatroom

	// err := json.Unmarshal(data, &decodedData)
	// if err != nil {
	// 	return nil, errors.New("Failed to process data")
	// }

	rows, err := query("GetChatrooms")
	if err != nil {
		return nil, errors.New("Failed to query chatrooms")
	}
	defer rows.Close()

	var chatrooms []Chatroom

	for rows.Next() {
		var chatroom Chatroom
		err := rows.Scan(&chatroom.ID, &chatroom.Title)
		if err != nil {
			return nil, errors.New("Failed to scan rows")
		}
		chatrooms = append(chatrooms, chatroom)
	}

	response, err := json.Marshal(map[string]interface{}{
		"status":    "Success",
		"chatrooms": chatrooms,
	})
	if err != nil {
		return nil, errors.New("cannot marshal chatrooms")
	}

	return response, nil
}
