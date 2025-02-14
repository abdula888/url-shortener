package postgres

import (
	"database/sql"
	"log"
	"url-shortener/pkg/response"

	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetURL(alias string) (string, error) {
	stmt, err := r.db.Prepare("SELECT url FROM url WHERE alias = $1")
	if err != nil {
		return "", err
	}
	var resURL string

	err = stmt.QueryRow(alias).Scan(&resURL)
	if err != nil {
		return "", response.ErrURLNotFound
	}
	return resURL, nil
}
func (r *Repository) SaveURL(alias, url string) error {
	stmt, err := r.db.Prepare("INSERT INTO url(alias, url) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return err
	}

	err = stmt.QueryRow(alias, url).Err()
	if err != nil {
		// Проверка ошибки на уникальность поля
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return response.ErrURLExists
			}
		}
		return err
	}
	return nil
}
