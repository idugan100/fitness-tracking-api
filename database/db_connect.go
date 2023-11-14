package database

import (
	"database/sql"
	"log"
)

var DB_connection *sql.DB

func Connect_to_database() {
	var err error
	DB_connection, err = sql.Open("sqlite3", "file:./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to databse sucessfully")
}
