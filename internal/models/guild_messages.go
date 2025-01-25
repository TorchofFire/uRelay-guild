package models

type GuildMessages struct {
	ID        uint64 `json:"id" db:"id"`
	SenderID  uint64 `json:"sender_id" db:"sender_id"`
	Message   string `json:"message" db:"message"`
	ChannelID uint64 `json:"channel_id" db:"channel_id"`
	SentAt    uint32 `json:"sent_at" db:"sent_at"`
}
