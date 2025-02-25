package usecase

import (
	"time"
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/repository"

	"math/rand"
)

type Usecase struct {
	repository repository.Repository
}

func New(repo repository.Repository) *Usecase {
	return &Usecase{
		repo,
	}
}

func (u *Usecase) GetURL(alias string) (string, error) {
	urlEntity, err := u.repository.GetURL(alias)
	if err != nil {
		return "", err
	}
	return urlEntity.OriginalURL, nil
}

func (u *Usecase) SaveURL(originalURL string) (string, error) {
	alias := newRandomString(10)
	urlEntity := entity.URL{Alias: alias, OriginalURL: originalURL}
	err := u.repository.SaveURL(urlEntity)
	if err != nil {
		return "", err
	}

	return alias, nil
}

func newRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
