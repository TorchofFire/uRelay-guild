package connections

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/TorchofFire/uRelay-guild/internal/packets"
	"github.com/TorchofFire/uRelay-guild/internal/types"
	"github.com/gorilla/websocket"
)

func sendPacketToAll(packet packets.BasePacket) {
	packetJSON, err := json.Marshal(packet)
	if err != nil {
		log.Println("Failed to marshal packet:", err)
	}

	MapMu.Lock()
	defer MapMu.Unlock()

	var waitGroup sync.WaitGroup
	for _, conn := range Map {
		waitGroup.Add(1)
		go func(c *websocket.Conn) {
			defer waitGroup.Done()
			c.WriteMessage(websocket.TextMessage, packetJSON)
		}(conn)
	}

	waitGroup.Wait()
}

func sendSystemMessageViaConn(conn *websocket.Conn, severity types.Severity, message string, channelId int) {
	sysMessage := packets.SystemMessage{
		Severity:  severity,
		Message:   message,
		ChannelId: channelId,
	}
	sysMessageJSON, err := json.Marshal(sysMessage)
	if err != nil {
		log.Println("could not encode system packet into json:", err)
	}
	packet := packets.BasePacket{
		Type: types.SystemMessage,
		Data: sysMessageJSON,
	}
	if err := conn.WriteJSON(packet); err != nil {
		conn.Close()
	}
}
