package handler

import (
	"encoding/json"
	"errors"
)

// This should include a token and validated server side (currently allows user id spoofing to create messages as another user)
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

	var username string
	//GetUsernameByID
	rows, err := query("GetUsernameByID", decodedData.UserId)
	if err != nil || !rows.Next() {
		return nil, errors.New("Failed to get username")
	}
	rows.Scan(&username)

	var dataMap map[string]interface{}
	json.Unmarshal(data, &dataMap)
	dataMap["username"] = username
	data, err = json.Marshal(dataMap)
	if err != nil {
		return nil, errors.New("Failed to add username to message data")
	}

	chatroomManager.SendMessageToChatroom(decodedData.ChatroomId, data)

	return json.Marshal(map[string]interface{}{
		"status":     "Success",
		"message":    decodedData,
		"chatroomID": decodedData.ChatroomId,
	})
}
