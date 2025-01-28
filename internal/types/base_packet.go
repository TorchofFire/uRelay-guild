package types

type BasePacket string

const (
	Handshake     BasePacket = "handshake"
	GuildMessage  BasePacket = "guild_message"
	SystemMessage BasePacket = "system_message"
	User          BasePacket = "user"
)
