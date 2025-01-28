package routes

import (
	"encoding/json"
	"net/http"
)

type userStripped struct {
	ID        uint64 `json:"id"`
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
	Status    status `json:"status"`
}

type status string

const (
	Online status = "online"
	// Idle    status = "idle"
	// DnD     status = "dnd"
	Offline status = "offline"
)

func (s *Service) users(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []userStripped
	for _, user := range s.guild.GetUsers() {
		userStatus := Offline
		if s.connections.UserConnected(user.ID) {
			userStatus = Online
		}
		users = append(users, userStripped{
			ID:        user.ID,
			PublicKey: user.PublicKey,
			Name:      user.Name,
			Status:    userStatus,
		})
	}
	if len(users) == 0 {
		users = []userStripped{}
	}
	if err := json.NewEncoder(writer).Encode(users); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
