package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/TorchofFire/uRelay-guild/internal/database"
	"github.com/TorchofFire/uRelay-guild/internal/models"
	"github.com/gorilla/mux"
)

func textChannel(writer http.ResponseWriter, request *http.Request) { // TODO: add perms to check if user can GET for this *specific* channel id
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	stringChannelId := vars["id"]
	channelId, err := strconv.Atoi(stringChannelId)
	if err != nil {
		http.Error(writer, "channel id must be a number", http.StatusBadRequest)
		return
	}

	var messages []models.GuildMessages
	var messagesQuery strings.Builder
	var queryArgs []interface{}

	messagesQuery.WriteString("SELECT * FROM guild_messages WHERE channel_id = ?")
	queryArgs = append(queryArgs, channelId)

	stringMsgId := request.URL.Query().Get("msg")
	if stringMsgId != "" {
		msgId, err := strconv.Atoi(stringMsgId)
		if err != nil {
			http.Error(writer, "msg param (message id) must be a number", http.StatusBadRequest)
			return
		}
		messagesQuery.WriteString(" AND id >= ?")
		queryArgs = append(queryArgs, msgId)
	}
	messagesQuery.WriteString(" ORDER BY id ASC LIMIT 15")

	result, err := database.DB.Query(messagesQuery.String(), queryArgs...)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()

	for result.Next() {
		var message models.GuildMessages
		if err := result.Scan(&message.ID, &message.SenderID, &message.Message, &message.ChannelID, &message.SentAt); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}

	if err := result.Err(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []models.GuildMessages{}
	}
	if err := json.NewEncoder(writer).Encode(messages); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
