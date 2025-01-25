package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TorchofFire/uRelay-guild/internal/models"
	"github.com/gorilla/mux"
)

func (s *Service) profile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	stringIds := vars["id"]
	userIds := strings.Split(stringIds, ",")
	if len(userIds) > 15 {
		userIds = userIds[:15]
	}

	userIdMap := make(map[uint64]bool)
	for _, id := range userIds {
		var idInt uint64
		_, err := fmt.Sscanf(id, "%d", &idInt)
		if err == nil {
			userIdMap[idInt] = true
		}
	}

	var profiles []models.Users
	for _, user := range s.guild.GetUsers() {
		if userIdMap[user.ID] {
			profiles = append(profiles, user)
		}
	}

	if profiles == nil {
		profiles = []models.Users{}
	}
	if err := json.NewEncoder(writer).Encode(profiles); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
