package database

import (
	"database/sql"
	"log"
)

func ConnectToDatabase() (*sql.DB, error) {
	DB_connection, err := sql.Open("sqlite3", "file:./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to databse sucessfully")
	return DB_connection, nil
}
