package types

type Status string

const (
	Online Status = "online"
	// Idle    status = "idle"
	// DnD     status = "dnd"
	Offline Status = "offline"
)
