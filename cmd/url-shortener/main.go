package main

import (
	"log"
	"url-shortener/internal/app"
	"url-shortener/internal/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Run(conf)
}
