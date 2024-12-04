package connections

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Map   = make(map[int]*websocket.Conn)
	MapMu sync.Mutex
)

func addNewConnection(userId int, conn *websocket.Conn) {
	MapMu.Lock()
	Map[userId] = conn
	MapMu.Unlock()
	log.Printf("User %d connected", userId)
}

func removeConnection(userId int) {
	MapMu.Lock()
	delete(Map, userId)
	MapMu.Unlock()
	log.Printf("User %d disconnected", userId)
}
