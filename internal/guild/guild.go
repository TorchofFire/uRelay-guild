package guild

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/TorchofFire/uRelay-guild/internal/models"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db         *sqlx.DB
	users      []models.Users
	channels   []models.GuildChannels
	categories []models.GuildCategories
}

func NewService(db *sqlx.DB) *Service {
	s := &Service{db: db}
	s.updateUsersCache()
	s.updateChannelsCache()
	s.updateCategoriesCache()
	return s
}

func (s *Service) updateUsersCache() {
	const usersQuery = "SELECT * FROM users;"

	err := s.db.Select(&s.users, usersQuery)
	if err != nil {
		log.Fatalf("Error updating users cache: %v", err)
	}
	if len(s.users) == 0 {
		log.Println("Database fetch successful but it appears there are no users!")
	}
}

func (s *Service) updateChannelsCache() {
	const channelsQuery = "SELECT * FROM guild_channels;"

	err := s.db.Select(&s.channels, channelsQuery)
	if err != nil {
		log.Fatalf("Error updating channels cache: %v", err)
	}
	if len(s.channels) == 0 {
		log.Println("Database fetch successful but it appears there are no channels!")
	}
}

func (s *Service) updateCategoriesCache() {
	const categoriesQuery = "SELECT * FROM guild_categories"

	err := s.db.Select(&s.categories, categoriesQuery)
	if err != nil {
		log.Fatalf("Error updating categories cache: %v", err)
	}
	if len(s.categories) == 0 {
		log.Println("Database fetch successful but it appears there are no categories!")
	}
}

func (s *Service) GetUsers() []models.Users {
	return s.users
}

func (s *Service) GetChannels() []models.GuildChannels {
	return s.channels
}

func (s *Service) GetCategories() []models.GuildCategories {
	return s.categories
}

func (s *Service) GetUser(id uint64) (models.Users, error) {
	var user models.Users
	for _, u := range s.GetUsers() {
		if u.ID == id {
			user = u
			break
		}
	}
	if user.ID == 0 {
		return models.Users{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Service) GetGuildMessages(channelId, msgId int) ([]models.GuildMessages, error) {
	var messages []models.GuildMessages
	var messagesQuery strings.Builder
	var queryArgs []interface{}

	messagesQuery.WriteString("SELECT * FROM guild_messages WHERE channel_id = ?")
	queryArgs = append(queryArgs, channelId)

	if msgId != 0 {
		messagesQuery.WriteString(" AND id <= ?")
		queryArgs = append(queryArgs, msgId)
	}
	messagesQuery.WriteString(" ORDER BY id DESC LIMIT 15")

	result, err := s.db.Query(messagesQuery.String(), queryArgs...)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer result.Close()

	for result.Next() {
		var message models.GuildMessages
		if err := result.Scan(&message.ID, &message.SenderID, &message.Message, &message.ChannelID, &message.SentAt); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		messages = append(messages, message)
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if messages == nil {
		messages = []models.GuildMessages{}
	}

	return messages, nil
}

func (s *Service) AddNewUser(publicKey, name string) (uint64, error) {
	const userInsert = "INSERT INTO users (public_key, name) VALUES (?, ?);"

	result, err := s.db.Exec(userInsert, publicKey, name)
	if err != nil {
		return 0, fmt.Errorf("failed to insert new user: %w", err)
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve new user ID: %w", err)
	}

	newUser := models.Users{
		ID:        uint64(newUserId),
		PublicKey: publicKey,
		Name:      name,
		JoinDate:  uint32(time.Now().Unix()),
	}
	s.users = append(s.users, newUser)

	return uint64(newUserId), nil
}

func (s *Service) InsertGuildMessage(senderId, channelId uint64, message string) (uint64, error) {
	const guildMessageInsert = "INSERT INTO guild_messages (sender_id, message, channel_id) VALUES (?, ?, ?);"
	result, err := s.db.Exec(guildMessageInsert, senderId, message, channelId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert message: %v", err)
	}
	messageId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve message id: %v", err)
	}
	return uint64(messageId), err
}
