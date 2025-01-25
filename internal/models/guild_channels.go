package models

import (
	"github.com/TorchofFire/uRelay-guild/internal/types"
)

type GuildChannels struct {
	ID              uint64            `json:"id" db:"id"`
	CategoryID      *uint64           `json:"category_id" db:"category_id"`
	Name            string            `json:"name" db:"name"`
	ChannelType     types.ChannelType `json:"channel_type" db:"channel_type"`
	DisplayPriority uint16            `json:"display_priority" db:"display_priority"`
}
