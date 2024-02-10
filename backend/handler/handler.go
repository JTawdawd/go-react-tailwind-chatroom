package handler

import "chatroom"

var chatroomManager *chatroom.ChatroomManager

func SetChatroomManager(cm *chatroom.ChatroomManager) {
	chatroomManager = cm
}
