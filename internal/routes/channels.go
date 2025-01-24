package routes

import (
	"encoding/json"
	"net/http"
)

func (s *Service) channels(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(s.guild.GetChannels()); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
