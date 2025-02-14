package inmem

import (
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

func (r *Repository) GetURL(alias string) (string, error) {
	url, exists := r.urls[alias]
	if !exists {
		return "", response.ErrURLNotFound
	}

	return url, nil
}
func (r *Repository) SaveURL(alias, url string) error {
	if _, exists := r.urls[alias]; exists {
		return response.ErrURLExists
	}

	r.urls[alias] = url

	return nil
}
