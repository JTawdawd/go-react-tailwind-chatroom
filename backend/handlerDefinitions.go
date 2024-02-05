package main

import (
	"handler"
)

type handlerDefinition func([]byte) ([]byte, error)

var handlerDefinitions map[string]handlerDefinition

func init() {
	handlerDefinitions = make(map[string]handlerDefinition)

	//handlerDefinitions["/"] = fs
	handlerDefinitions["/api/login"] = handler.HandleLogin
	handlerDefinitions["/api/createUser"] = handler.CreateUser
	handlerDefinitions["/api/getChatrooms"] = handler.GetChatrooms
	handlerDefinitions["/api/getChatroom"] = handler.GetChatroom
	handlerDefinitions["/api/createMessage"] = handler.CreateMessage
	handlerDefinitions["/api/createChatroom"] = handler.CreateChatroom
}
