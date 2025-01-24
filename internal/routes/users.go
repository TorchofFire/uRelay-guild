package routes

import (
	"encoding/json"
	"net/http"
)

type userStripped struct {
	ID        int    `json:"id"`
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
}

func (s *Service) users(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []userStripped
	for _, user := range s.guild.GetUsers() {
		// TODO: if id exists in connections, set property online
		users = append(users, userStripped{
			ID:        user.ID,
			PublicKey: user.PublicKey,
			Name:      user.Name,
		})
	}
	if err := json.NewEncoder(writer).Encode(users); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
