package types

type BasePacket string

const (
	handshake     BasePacket = "handshake"
	GuildMessage  BasePacket = "guild_message"
	SystemMessage BasePacket = "system_message"
)
