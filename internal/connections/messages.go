package connections

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/TorchofFire/uRelay-guild/config"
	"github.com/TorchofFire/uRelay-guild/internal/database"
	"github.com/TorchofFire/uRelay-guild/internal/guild"
	"github.com/TorchofFire/uRelay-guild/internal/models"
	"github.com/TorchofFire/uRelay-guild/internal/packets"
	"github.com/TorchofFire/uRelay-guild/internal/types"
	"github.com/gorilla/websocket"
)

func handshake(packet packets.Handshake) (int, error) {
	serverId, err := unlockAndVerifySignedMessage(packet.PublicKey, packet.Proof, 30)
	if err != nil {
		return 0, err
	}
	if serverId != config.ServerID {
		return 0, fmt.Errorf("expected server identifer; looking for %s, instead got %s; format is timestamp|serverid", config.ServerID, serverId)
	}

	for _, user := range guild.Users {
		if user.PublicKey == packet.PublicKey {
			return user.ID, nil
		}
	}

	// New user. Add to DB and give resulting id.

	const userInsert = "INSERT INTO users (public_key, name) VALUES (?, ?);"

	result, err := database.DB.Exec(userInsert, packet.PublicKey, packet.Name)
	if err != nil {
		return 0, fmt.Errorf("failed to insert new user: %w", err)
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve new user ID: %w", err)
	}

	newUser := models.Users{
		ID:        int(newUserId),
		PublicKey: packet.PublicKey,
		Name:      packet.Name,
		JoinDate:  int(time.Now().Unix()),
	}
	guild.Users = append(guild.Users, newUser)

	return int(newUserId), nil
}

func handleGuildMessage(conn *websocket.Conn, packet packets.GuildMessage) {
	var channel *models.GuildChannels
	for _, ch := range guild.Channels {
		if ch.ID == packet.ChannelId {
			channel = &ch
			break
		}
	}
	if channel == nil {
		sendSystemMessageViaConn(conn, types.Danger, "Channel not found", packet.ChannelId)
		return
	}

	var user *models.Users
	for _, u := range guild.Users {
		if u.ID == packet.SenderId {
			user = &u
			break
		}
	}
	if user == nil {
		sendSystemMessageViaConn(conn, types.Danger, "User not found", packet.ChannelId)
		return
	}

	_, err := unlockAndVerifySignedMessage(user.PublicKey, packet.Message, 30)
	if err != nil {
		sendSystemMessageViaConn(conn, types.Danger, "Your message didn't satisfy security requirements: "+err.Error(), packet.ChannelId)
		return
	}

	// TODO: verify user has perms to send in specific channel

	const guildMessageInsert = "INSERT INTO guild_messages (sender_id, message, channel_id) VALUES (?, ?, ?);"
	result, err := database.DB.Exec(guildMessageInsert, user.ID, packet.Message, packet.ChannelId)
	if err != nil {
		log.Println("failed to insert message:", err)
		sendSystemMessageViaConn(conn, types.Danger, "Server failure: "+err.Error(), packet.ChannelId)
		return
	}
	messageId, err := result.LastInsertId()
	if err != nil {
		log.Println("failed to retrieve message ID:", err)
		return
	}

	packet.Id = int(messageId)
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
