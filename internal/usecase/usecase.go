package usecase

import "url-shortener/internal/lib/random"

type repository interface {
	GetURL(alias string) (string, error)
	SaveURL(alias, url string) (bool, error)
}

type Usecase struct {
	repository
}

func New(repo repository) *Usecase {
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
	alias := random.NewRandomString(10)
	aliasExists, err := u.repository.SaveURL(alias, url)
	if aliasExists { // Если генератор ссылок выдал ссылку, которая уже есть в БД, пытаемся снова
		u.SaveURL(url)
	}
	if err != nil {
		return "", err
	}

	return alias, nil
}
