package packets

import (
	"encoding/json"

	"github.com/TorchofFire/uRelay-guild/internal/types"
)

type Service struct {
}

func NewService() *Service {
	s := &Service{}
	return s
}

type BasePacket struct {
	Type types.BasePacket `json:"type"`
	Data json.RawMessage  `json:"data"`
}

type Handshake struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
	Proof     string `json:"proof"`
	/*
		proof is timestamp with server id signed.
		example of decoded: unixtimestamp|ipOrDomain(:port)
		"|" being a delimiter.
		The server id is important so the handshake is server specific and cannot be replayed elsewhere.
	*/
}

type GuildMessage struct {
	ChannelId uint64 `json:"channel_id"`
	SenderId  uint64 `json:"sender_id"`
	Message   string `json:"message"`
	Id        uint64 `json:"id"`
}

type SystemMessage struct {
	Severity  types.Severity `json:"severity"`
	Message   string         `json:"message"`
	ChannelId uint64         `json:"channel_id"`
}

type User struct {
	ID        uint64       `json:"id"`
	PublicKey string       `json:"public_key"`
	Name      string       `json:"name"`
	Status    types.Status `json:"status"`
}
