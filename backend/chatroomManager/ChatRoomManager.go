package chatroom

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ChatroomManager struct {
	connections map[string][]*websocket.Conn
	mutexes     map[string]*sync.Mutex
}

func CreateChatroom() *ChatroomManager {
	return &ChatroomManager{
		connections: make(map[string][]*websocket.Conn),
		mutexes:     make(map[string]*sync.Mutex),
	}
}

func (cm *ChatroomManager) AddConnectionToChatroom(chatroomID string, connection *websocket.Conn) {
	cm.getMutex(chatroomID).Lock()
	defer cm.getMutex(chatroomID).Unlock()

	if _, ok := cm.connections[chatroomID]; !ok {
		cm.connections[chatroomID] = make([]*websocket.Conn, 0)
	}

	cm.connections[chatroomID] = append(cm.connections[chatroomID], connection)
}

func (cm *ChatroomManager) SendMessageToChatroom(chatroomID string, message []byte) {
	cm.getMutex(chatroomID).Lock()
	defer cm.getMutex(chatroomID).Unlock()

	connections, ok := cm.connections[chatroomID]
	if !ok {
		return
	}

	for _, connection := range connections {
		connection.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func (cm *ChatroomManager) RemoveConnectionFromChatroom(chatroomID string, connection *websocket.Conn) {
	cm.getMutex(chatroomID).Lock()
	defer cm.getMutex(chatroomID).Unlock()

	for i, c := range cm.connections[chatroomID] {
		if c == connection {
			cm.connections[chatroomID] = append(cm.connections[chatroomID][:i], cm.connections[chatroomID][i+1:]...)
			break
		}
	}
}

func (cm *ChatroomManager) getMutex(chatroomID string) *sync.Mutex {
	cm.mutexes[chatroomID] = cm.mutexes[chatroomID]
	if cm.mutexes[chatroomID] == nil {
		cm.mutexes[chatroomID] = &sync.Mutex{}
	}
	return cm.mutexes[chatroomID]
}
