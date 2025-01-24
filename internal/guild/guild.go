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
	db       *sqlx.DB
	users    []models.Users
	channels []models.GuildChannels
}

func NewService(db *sqlx.DB) *Service {
	s := &Service{db: db}
	s.updateUsersCache()
	s.updateChannelsCache()
	return s
}

func (s *Service) updateUsersCache() {
	const usersQuery = "SELECT id, public_key, name, join_date FROM users;"

	err := s.db.Select(&s.users, usersQuery)
	if err != nil {
		log.Fatalf("Error updating users cache: %v", err)
	}
	if len(s.users) == 0 {
		log.Println("Database fetch successful but it appears there are no users!")
	}
}

func (s *Service) updateChannelsCache() {
	const channelsQuery = "SELECT id, name, channel_type FROM guild_channels;"

	err := s.db.Select(&s.channels, channelsQuery)
	if err != nil {
		log.Fatalf("Error updating channels cache: %v", err)
	}
	if len(s.channels) == 0 {
		log.Println("Database fetch successful but it appears there are no channels!")
	}
}

func (s *Service) GetChannels() []models.GuildChannels {
	return s.channels
}

func (s *Service) GetUsers() []models.Users {
	return s.users
}

func (s *Service) GetUser(id int) (models.Users, error) {
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
		messagesQuery.WriteString(" AND id >= ?")
		queryArgs = append(queryArgs, msgId)
	}
	messagesQuery.WriteString(" ORDER BY id ASC LIMIT 15")

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

func (s *Service) AddNewUser(publicKey, name string) (int, error) {
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
		ID:        int(newUserId),
		PublicKey: publicKey,
		Name:      name,
		JoinDate:  int(time.Now().Unix()),
	}
	s.users = append(s.users, newUser)

	return int(newUserId), nil
}

func (s *Service) InsertGuildMessage(senderId, channelId int, message string) (int, error) {
	const guildMessageInsert = "INSERT INTO guild_messages (sender_id, message, channel_id) VALUES (?, ?, ?);"
	result, err := s.db.Exec(guildMessageInsert, senderId, message, channelId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert message: %v", err)
	}
	messageId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve message id: %v", err)
	}
	return int(messageId), err
}
