package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DatabaseURL string
		LogLevel    string
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load("config/.env")

	databaseURL := os.Getenv("DATABASE_URL")
	logLevel := os.Getenv("LOG_LEVEL")

	if err != nil {
		return nil, err
	}

	conf := &Config{DatabaseURL: databaseURL, LogLevel: logLevel}

	return conf, nil
}
