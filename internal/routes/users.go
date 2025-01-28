package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/internal/packets"
	"github.com/TorchofFire/uRelay-guild/internal/types"
)

func (s *Service) users(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []packets.User
	for _, user := range s.guild.GetUsers() {
		userStatus := types.Offline
		if s.connections.UserConnected(user.ID) {
			userStatus = types.Online
		}
		users = append(users, packets.User{
			ID:        user.ID,
			PublicKey: user.PublicKey,
			Name:      user.Name,
			Status:    userStatus,
		})
	}
	if len(users) == 0 {
		users = []packets.User{}
	}
	if err := json.NewEncoder(writer).Encode(users); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
