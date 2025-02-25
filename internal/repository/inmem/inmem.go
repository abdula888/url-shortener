package inmem

import (
	"url-shortener/internal/domain/entity"
	"url-shortener/pkg/response"
)

type Repository struct {
	urls map[string]string
}

func NewRepo() *Repository {
	return &Repository{
		urls: make(map[string]string),
	}
}

func (r *Repository) GetURL(alias string) (entity.URL, error) {
	originalURL, exists := r.urls[alias]
	if !exists {
		return entity.URL{}, response.ErrURLNotFound
	}
	urlEntity := entity.URL{Alias: alias, OriginalURL: originalURL}

	return urlEntity, nil
}
func (r *Repository) SaveURL(urlEntity entity.URL) error {
	if _, exists := r.urls[urlEntity.Alias]; exists {
		return response.ErrURLExists
	}

	r.urls[urlEntity.Alias] = urlEntity.OriginalURL

	return nil
}
