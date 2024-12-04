package models

type GuildMessages struct {
	ID        int    `json:"id" db:"id"`
	SenderID  string `json:"sender_id" db:"sender_id"`
	Message   string `json:"message" db:"message"`
	ChannelID int    `json:"channel_id" db:"channel_id"`
	SentAt    int    `json:"sent_at" db:"sent_at"`
}
