package database

import (
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

type HandlerDB struct {
	DB *sql.DB
}

var Handler HandlerDB

func Init() (HandlerDB, error) {
	file, err := os.Create("bitly.db")
	if err != nil {
		return HandlerDB{}, err
	}
	defer file.Close()
	
	//open database
	db, err := sql.Open("sqlite3", "bitly.db")
	if err != nil {
		return HandlerDB{}, err
	}

	//sql code to create table
	createTable := `
		CREATE TABLE bitly (
			id VARCHAR(7),
			url VARCHAR(255)
		);
	`

	//execute query
	if _, err := db.Exec(createTable); err != nil {
		return HandlerDB{}, err
	}

	return HandlerDB{db}, nil
}