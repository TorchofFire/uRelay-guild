package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TorchofFire/uRelay-guild/internal/guild"
	"github.com/TorchofFire/uRelay-guild/internal/models"
)

func profile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(request.URL.Path, "/profile/")
	userIds := strings.Split(path, ",")
	if len(userIds) > 15 {
		userIds = userIds[:15]
	}

	userIdMap := make(map[int]bool)
	for _, id := range userIds {
		var idInt int
		_, err := fmt.Sscanf(id, "%d", &idInt)
		if err == nil {
			userIdMap[idInt] = true
		}
	}

	var profiles []models.Users
	for _, user := range guild.Users {
		if userIdMap[user.ID] {
			profiles = append(profiles, user)
		}
	}

	if len(profiles) == 0 {
		http.NotFound(writer, request)
		return
	}

	if err := json.NewEncoder(writer).Encode(profiles); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
