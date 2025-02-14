package usecase

import (
	"time"
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
	url, err := u.repository.GetURL(alias)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *Usecase) SaveURL(url string) (string, error) {
	alias := newRandomString(10)
	err := u.repository.SaveURL(alias, url)
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
