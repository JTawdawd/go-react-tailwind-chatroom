package handler

import (
	"encoding/json"
	"errors"
)

type chatroomRequest2 struct {
	Title string `json:"title"`
}

func CreateChatroom(data []byte) ([]byte, error) {
	var decodedData chatroomRequest2
	err := json.Unmarshal(data, &decodedData)
	if err != nil {
		return nil, err
	}

	rows, err := query("ChatroomByTitle", decodedData.Title)
	if err != nil {
		return nil, errors.New("Failed to check for existing chatrooms")
	}
	if rows.Next() {
		return nil, errors.New("Chatroom already exists")
	}

	err = insert("CreateChatroom", decodedData.Title)
	if err != nil {
		return nil, errors.New("Failed to create chatroom")
	}

	var (
		id int
	)

	rows, err = query("ChatroomByTitle", decodedData.Title)
	if err != nil || !rows.Next() {
		return nil, errors.New("Failed to get created chatroom")
	}
	rows.Scan(&id)

	return json.Marshal(map[string]interface{}{
		"status": "Success",
		"title":  decodedData.Title,
		"id":     id,
	})
}
