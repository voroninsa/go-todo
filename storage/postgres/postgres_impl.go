package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *postgresStorage) GetUrl(shortUrl string) (string, error) {
	var url string

	if err := s.StorageURL.QueryRow(`SELECT url FROM urls WHERE short_url = $1`, shortUrl).Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New(fmt.Sprint("No such record: ", shortUrl))
		}
		return "", errors.New(fmt.Sprint("Error DataBase: ", err))
	}

	return url, nil
}

func (s *postgresStorage) PutUrl(shortUrl string, url string) error {
	db_url, _ := s.GetUrl(shortUrl)
	// Check if record exists. Don't check errors in Get
	if db_url != "" {
		return nil
	}

	_, err := s.StorageURL.Exec(`INSERT INTO urls (short_url, url) VALUES ($1, $2)`, shortUrl, url)
	if err != nil {
		return errors.New(fmt.Sprint("Adding data error: ", err))
	}

	return nil
}
