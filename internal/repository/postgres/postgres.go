package postgres

import (
	"database/sql"
	"log"
	"url-shortener/internal/domain/entity"
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

func (r *Repository) GetURL(alias string) (entity.URL, error) {
	stmt, err := r.db.Prepare("SELECT url FROM url WHERE alias = $1")
	if err != nil {
		return entity.URL{}, err
	}
	var originalURL string

	err = stmt.QueryRow(alias).Scan(&originalURL)
	if err != nil {
		return entity.URL{}, response.ErrURLNotFound
	}
	urlEntity := entity.URL{Alias: alias, OriginalURL: originalURL}

	return urlEntity, nil
}
func (r *Repository) SaveURL(urlEntity entity.URL) error {
	stmt, err := r.db.Prepare("INSERT INTO url(alias, url) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return err
	}

	err = stmt.QueryRow(urlEntity.Alias, urlEntity.OriginalURL).Err()
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
