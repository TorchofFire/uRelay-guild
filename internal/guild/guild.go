package guild

import (
	"log"

	"github.com/TorchofFire/uRelay-guild/internal/database"
	"github.com/TorchofFire/uRelay-guild/internal/models"
)

var Users []models.Users
var Channels []models.GuildChannels

func Init() {
	updateUsersCache()
	updateChannelsCache()
}

func updateUsersCache() {
	const usersQuery = "SELECT id, public_key, name, join_date FROM users;"

	err := database.DB.Select(&Users, usersQuery)
	if err != nil {
		log.Fatalf("Error updating users cache: %v", err)
	}
	if len(Users) == 0 {
		log.Println("Database fetch successful but it appears there are no users!")
	}
}

func updateChannelsCache() {
	const channelsQuery = "SELECT id, name, channel_type FROM guild_channels;"

	err := database.DB.Select(&Channels, channelsQuery)
	if err != nil {
		log.Fatalf("Error updating channels cache: %v", err)
	}
	if len(Users) == 0 {
		log.Println("Database fetch successful but it appears there are no channels!")
	}
}
