package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/internal/models"
)

type channelsAndCategories struct {
	Channels   []models.GuildChannels   `json:"channels"`
	Categories []models.GuildCategories `json:"categories"`
}

func (s *Service) channels(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := channelsAndCategories{
		Channels:   s.guild.GetChannels(),
		Categories: s.guild.GetCategories(),
	}
	if len(data.Channels) == 0 {
		data.Channels = []models.GuildChannels{}
	}
	if len(data.Categories) == 0 {
		data.Categories = []models.GuildCategories{}
	}
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
