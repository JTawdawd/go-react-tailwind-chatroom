package main

import (
	"handler"
)

type handlerDefinition func([]byte) ([]byte, error)

var handlerDefinitions map[string]handlerDefinition

func init() {
	handlerDefinitions = make(map[string]handlerDefinition)

	//handlerDefinitions["/"] = fs
	handlerDefinitions["/login"] = handler.HandleLogin
	handlerDefinitions["/create"] = handler.CreateUser
	handlerDefinitions["chatrooms"] = handler.GetChatrooms
	handlerDefinitions["/chatroom"] = handler.GetChatroom
}
