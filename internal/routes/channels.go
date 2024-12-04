package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TorchofFire/uRelay-guild/internal/guild"
)

func channels(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(guild.Channels); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
