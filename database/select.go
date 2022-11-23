package database

import (
	_ "github.com/mattn/go-sqlite3"
)

//select url from database by id
func Select(id string) (string, error) {
	//get url
	row := Handler.DB.QueryRow("SELECT url FROM bitly WHERE id=?;", id)
	var url string
	if err := row.Scan(&url); err != nil {
		return "", nil
	}
	
	return url, nil
}