package models

type Users struct {
	ID        int    `json:"id" db:"id"`
	PublicKey string `json:"public_key" db:"public_key"`
	Name      string `json:"name" db:"name"`
	JoinDate  int    `json:"join_date" db:"join_date"`
}
