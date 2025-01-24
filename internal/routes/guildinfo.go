package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/internal/connections"
)

type gInfo struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Logo            string `json:"logo"`
	Banner          string `json:"banner"`
	UserCount       int    `json:"user_count"`
	OnlineUserCount int    `json:"online_user_count"`
}

func (s *Service) guildInfo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	guildInfo := gInfo{
		Name:            "", // TODO: add guild name, version and image links
		Version:         "",
		Logo:            "",
		Banner:          "",
		UserCount:       len(s.guild.GetUsers()),
		OnlineUserCount: len(connections.Map),
	}
	if err := json.NewEncoder(writer).Encode(guildInfo); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
