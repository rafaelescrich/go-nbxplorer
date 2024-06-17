package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName          string
	Environment      string
	RabbitMQURL      string
	SentryDSN        string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	BTCRPCURL        string
	BTCNodePort      string
}

var AppConfig *Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, %s", err)
	}

	AppConfig = &Config{
		AppName:          getEnv("APP_NAME", "NBXplorer"),
		Environment:      getEnv("ENVIRONMENT", "development"),
		RabbitMQURL:      getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		SentryDSN:        getEnv("SENTRY_DSN", ""),
		PostgresUser:     getEnv("POSTGRES_USER", "user"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
		PostgresDB:       getEnv("POSTGRES_DB", "nbxplorer"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		BTCRPCURL:        getEnv("BTC_RPC_URL", "http://localhost"),
		BTCNodePort:      getEnv("BTC_NODE_PORT", "8332"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
