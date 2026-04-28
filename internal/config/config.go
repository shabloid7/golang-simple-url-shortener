package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	ServerPort string
	RedisAddr string
	BaseURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		BaseURL: os.Getenv("BASE_URL"),
	}
}