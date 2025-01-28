package connections

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/TorchofFire/uRelay-guild/internal/packets"
	"github.com/TorchofFire/uRelay-guild/internal/types"
	"github.com/gorilla/websocket"
)

func (s *Service) sendPacketToAll(packet packets.BasePacket) {
	MapMu.Lock()
	defer MapMu.Unlock()

	var waitGroup sync.WaitGroup
	for _, conn := range Map {
		waitGroup.Add(1)
		go func(c *websocket.Conn) {
			defer waitGroup.Done()
			c.WriteJSON(packet)
		}(conn)
	}

	waitGroup.Wait()
}

func (s *Service) sendUserToAll(userId uint64) {
	user, err := s.guild.GetUser(userId)
	if err != nil {
		log.Fatalf("tried to access user to send to all but failed: %v", err)
	}
	userStatus := types.Offline
	if s.UserConnected(user.ID) {
		userStatus = types.Online
	}
	packet := packets.User{
		ID:        user.ID,
		PublicKey: user.PublicKey,
		Name:      user.Name,
		Status:    userStatus,
	}
	packetJSON, err := json.Marshal(packet)
	if err != nil {
		log.Println("Failed to marshal packet:", err)
	}
	basePacket := packets.BasePacket{
		Type: types.User,
		Data: packetJSON,
	}
	s.sendPacketToAll(basePacket)
}

func (s *Service) sendSystemMessageViaConn(conn *websocket.Conn, severity types.Severity, message string, channelId uint64) {
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
