package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ServerID       string
	SecureProtocol bool
	CertPath       string
	DBHost         string
	DBUser         string
	DBPassword     string
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables:", err)
	}

	ServerID = getEnv("UR_GUILD_SERVER_ID", "localhost:8080")
	SecureProtocol = getEnvAsBool("UR_GUILD_SECURE_PROTOCOL", false)
	CertPath = getEnv("UR_GUILD_SSL_PATH", "")
	DBHost = getEnv("UR_GUILD_DB_HOST", "localhost")
	DBUser = getEnv("UR_GUILD_DB_USER", "root")
	DBPassword = getEnv("UR_GUILD_DB_PASSWORD", "")
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.ParseBool(valueStr)
		if err == nil {
			return value
		}
	}
	return defaultValue
}
