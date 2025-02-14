package app

import (
	"url-shortener/internal/config"
	"url-shortener/internal/delivery/http"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/inmem"
	"url-shortener/internal/repository/postgres"
	"url-shortener/internal/usecase"
	"url-shortener/pkg/db/conn"
	"url-shortener/pkg/db/migrate"
	"url-shortener/pkg/log"
)

func Run(conf *config.Config) {
	// Настройка уровня логирования
	log := log.NewLogger(conf.LogLevel)

	var repo repository.Repository

	if conf.StorageType == "inmem" {
		repo = inmem.NewRepo()
	} else {
		// Подключение к БД
		db, err := conn.InitDB()
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}
		log.Info("Successfully connected to the database")
		defer db.Close()

		// Применение миграций
		if err := migrate.RunMigrations(db); err != nil {
			log.Fatal("Error applying migration: ", err)
		}
		log.Info("Migrations applied successfully")
		repo = postgres.NewRepo(db)
	}

	// Регистрация маршрутов
	r := http.NewRouter(usecase.New(repo))
	log.Info("Routes registered successfully")

	// Запуск сервера
	log.Info("Server started on port 8080")
	log.Fatal(r.Run(":8080"))
}
