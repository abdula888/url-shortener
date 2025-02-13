package repository

import (
	"database/sql"
	"log"
	"url-shortener/internal/lib/response"
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetURL(alias string) (string, error) {
	// rows, err := r.db.Query("SELECT id, alias, url FROM url")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// // Итерация по строкам и вывод данных
	// for rows.Next() {
	// 	var id int
	// 	var alias, url string
	// 	err := rows.Scan(&id, &alias, &url)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("ID: %d, Alias: %s, URL: %s\n", id, alias, url)
	// }

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
func (r *Repo) SaveURL(alias, url string) (bool, error) {
	stmt, err := r.db.Prepare("INSERT INTO url(alias, url) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return false, err
	}

	err = stmt.QueryRow(alias, url).Err()
	if err != nil {
		log.Println(err)
		return false, err
	}
	return false, nil
}
