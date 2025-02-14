package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DatabaseURL  string
		DatabaseType string
		LogLevel     string
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load("config/.env")

	databaseURL := os.Getenv("DATABASE_URL")
	databaseType := os.Getenv("DATABASE_TYPE")
	logLevel := os.Getenv("LOG_LEVEL")

	if err != nil {
		return nil, err
	}

	conf := &Config{DatabaseURL: databaseURL, DatabaseType: databaseType, LogLevel: logLevel}

	return conf, nil
}
