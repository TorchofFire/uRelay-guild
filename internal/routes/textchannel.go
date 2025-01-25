package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Service) textChannel(writer http.ResponseWriter, request *http.Request) { // TODO: add perms to check if user can GET for this *specific* channel id
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	stringChannelId := vars["id"]
	channelId, err := strconv.Atoi(stringChannelId)
	if err != nil {
		http.Error(writer, "channel id must be a number", http.StatusBadRequest)
		return
	}

	stringMsgId := request.URL.Query().Get("msg")
	var msgId = 0
	if stringMsgId != "" {
		msgId, err = strconv.Atoi(stringMsgId)
		if err != nil {
			http.Error(writer, "msg param (message id) must be a number", http.StatusBadRequest)
			return
		}
	}

	messages, err := s.guild.GetGuildMessages(channelId, msgId)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(writer).Encode(messages); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
