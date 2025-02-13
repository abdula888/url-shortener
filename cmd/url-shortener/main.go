package main

import (
	"log"
	"url-shortener/internal/app"
	"url-shortener/internal/config"
)

// TODO: разобраться как правильно добавить логи в delivery, usecase и repo
func main() {
	// Инициализация конфига
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Run(conf)
}
