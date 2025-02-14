package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DatabaseURL string
		StorageType string
		LogLevel    string
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load("config/.env")

	databaseURL := os.Getenv("DATABASE_URL")
	storageType := os.Getenv("STORAGE_TYPE")
	logLevel := os.Getenv("LOG_LEVEL")

	if err != nil {
		return nil, err
	}

	conf := &Config{DatabaseURL: databaseURL, StorageType: storageType, LogLevel: logLevel}

	return conf, nil
}
