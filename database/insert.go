package database

import (
	_ "github.com/mattn/go-sqlite3"
)

//insert id and url into database
func Insert(id, url string) error {
	//sql code
	query := `
		INSERT INTO bitly(id, url) VALUES (?, ?);
	`

	//execute query
	if _, err := Handler.DB.Exec(query, id, url); err != nil {
		return err
	}

	return nil
}