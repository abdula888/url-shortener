package repository

import "url-shortener/internal/domain/entity"

type Repository interface {
	GetURL(alias string) (entity.URL, error)
	SaveURL(entity.URL) error
}
