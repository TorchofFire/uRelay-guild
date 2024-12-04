package models

import "github.com/TorchofFire/uRelay-guild/internal/types"

type GuildChannels struct {
	ID          int               `json:"id" db:"id"`
	Name        string            `json:"name" db:"name"`
	ChannelType types.ChannelType `json:"channel_type" db:"channel_type"`
}
