package models

type Users struct {
	ID        uint64 `json:"id" db:"id"`
	PublicKey string `json:"public_key" db:"public_key"`
	Name      string `json:"name" db:"name"`
	JoinDate  uint32 `json:"join_date" db:"join_date"`
}
