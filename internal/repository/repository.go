package repository

type Repository interface {
	GetURL(alias string) (string, error)
	SaveURL(alias, url string) error
}
