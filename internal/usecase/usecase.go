package usecase

type Usecase interface {
	GetURL(alias string) (string, error)
	SaveURL(url string) (string, error)
}
