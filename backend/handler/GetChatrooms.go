package handler

import (
	"encoding/json"
)

type Chatroom struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetChatrooms(data []byte) ([]byte, error) {

	var decodedData Chatroom

	err := json.Unmarshal(data, &decodedData)
	if err != nil {
		return nil, err
	}

	rows, err := query("GetChatrooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatrooms []Chatroom

	for rows.Next() {
		var chatroom Chatroom
		err := rows.Scan(&chatroom.ID, &chatroom.Title)
		if err != nil {
			return nil, err
		}
		chatrooms = append(chatrooms, chatroom)
	}

	return json.Marshal(chatrooms)
}
