package connections

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Map   = make(map[uint64]*websocket.Conn)
	MapMu sync.Mutex
)

// Note that it is assumed that a user is verified if added to the connections map.
// Since a user ID is provided to add a connection, then checks have already been ran.

func (s *Service) addNewConnection(userId uint64, conn *websocket.Conn) {
	s.sendUserToAll(userId)
	MapMu.Lock()
	Map[userId] = conn
	MapMu.Unlock()
	log.Printf("User %d connected", userId)
}

func (s *Service) removeConnection(userId uint64) {
	MapMu.Lock()
	delete(Map, userId)
	MapMu.Unlock()
	log.Printf("User %d disconnected", userId)
}

func (s *Service) UserConnected(userId uint64) bool {
	MapMu.Lock()
	defer MapMu.Unlock()
	_, exists := Map[userId]
	return exists
}
