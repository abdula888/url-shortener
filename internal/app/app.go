package app

import (
	"url-shortener/internal/config"
	"url-shortener/internal/delivery/http"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/inmem"
	"url-shortener/internal/repository/postgres"
	"url-shortener/internal/usecase/impl"
	"url-shortener/pkg/db/conn"
	"url-shortener/pkg/db/migrate"
	"url-shortener/pkg/log/logrus"
)

func Run(conf *config.Config) {
	// Настройка уровня логирования
	logger := logrus.NewLogger(conf.LogLevel)

	var repo repository.Repository

	if conf.StorageType == "inmem" {
		repo = inmem.NewRepo()
	} else {
		// Подключение к БД
		db, err := conn.InitDB(conf.DatabaseURL)
		if err != nil {
			logger.Fatal("Failed to connect to the database:", err)
		}
		logger.Info("Successfully connected to the database")
		defer db.Close()

		// Применение миграций
		if err := migrate.RunMigrations(db); err != nil {
			logger.Fatal("Error applying migration: ", err)
		}
		logger.Info("Migrations applied successfully")
		repo = postgres.NewRepo(db)
	}

	// Регистрация маршрутов
	r := http.NewRouter(impl.New(repo), logger)
	logger.Info("Routes registered successfully")

	// Запуск сервера
	logger.Info("Server started on port 8080")
	logger.Fatal(r.Run(":8080"))
}
