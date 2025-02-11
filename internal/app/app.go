package app

import (
	"url-shortener/internal/config"
	"url-shortener/pkg/db/conn"
	"url-shortener/pkg/db/migrate"
	"url-shortener/pkg/log"
)

func Run(conf *config.Config) {
	log := log.NewLogger(conf.LogLevel)

	log.Debug("debug messages are enabled")

	db, err := conn.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Info("Successfully connected to the database")
	defer db.Close()

	if err := migrate.RunMigrations(db); err != nil {
		log.Fatal("Error applying migration: ", err)
	}
	log.Info("Migrations applied successfully")
	//db.Exec("INSERT INTO url(url, alias) VALUES('52', 'spb')")
}
