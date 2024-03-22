package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDatabase() (*sql.DB, error) {
	DB_connection, err := sql.Open("sqlite3", "file:/Users/isaacdugan/code/fitness-tracker-api/database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to database sucessfully")
	return DB_connection, nil
}
