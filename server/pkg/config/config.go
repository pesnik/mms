package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	RabbitMQURL string
	RedisURL    string
	ServerPort  string
}

func New() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@postgres:5432/dbname"),
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672"),
		RedisURL:    getEnv("REDIS_URL", "redis://redis:6379"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
