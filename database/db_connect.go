package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDatabase(dbPath string) (*sql.DB, error) {
	DB_connection, err := sql.Open("sqlite3", ("file:" + dbPath))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to database sucessfully")
	return DB_connection, nil
}
