package connections

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TorchofFire/uRelay-guild/config"
	"github.com/TorchofFire/uRelay-guild/internal/models"
	"github.com/TorchofFire/uRelay-guild/internal/packets"
	"github.com/TorchofFire/uRelay-guild/internal/types"
	"github.com/gorilla/websocket"
)

func (s *Service) handshake(packet packets.Handshake) (uint64, error) {
	serverId, err := unlockAndVerifySignedMessage(packet.PublicKey, packet.Proof, 30)
	if err != nil {
		return 0, err
	}
	if serverId != config.ServerID {
		return 0, fmt.Errorf("expected server identifer; looking for %s, instead got %s; format is timestamp|serverid", config.ServerID, serverId)
	}

	for _, user := range s.guild.GetUsers() {
		if user.PublicKey == packet.PublicKey {
			return user.ID, nil
		}
	}

	return s.guild.AddNewUser(packet.PublicKey, packet.Name)
}

func (s *Service) handleGuildMessage(conn *websocket.Conn, packet packets.GuildMessage) {
	var channel *models.GuildChannels
	for _, ch := range s.guild.GetChannels() {
		if ch.ID == packet.ChannelId {
			channel = &ch
			break
		}
	}
	if channel == nil {
		sendSystemMessageViaConn(conn, types.Danger, "Channel not found", packet.ChannelId)
		return
	}

	user, err := s.guild.GetUser(packet.SenderId)
	if err != nil {
		sendSystemMessageViaConn(conn, types.Danger, err.Error(), packet.ChannelId)
		return
	}

	_, err = unlockAndVerifySignedMessage(user.PublicKey, packet.Message, 30)
	if err != nil {
		sendSystemMessageViaConn(conn, types.Danger, "Your message didn't satisfy security requirements: "+err.Error(), packet.ChannelId)
		return
	}

	// TODO: verify user has perms to send in specific channel

	packet.Id, err = s.guild.InsertGuildMessage(user.ID, packet.ChannelId, packet.Message)
	if err != nil {
		log.Println(err.Error())
		return
	}

	packetJSON, err := json.Marshal(packet)
	if err != nil {
		log.Println("Failed to marshal packet:", err)
	}
	basePacket := packets.BasePacket{
		Type: types.GuildMessage,
		Data: packetJSON,
	}

	sendPacketToAll(basePacket)
}
