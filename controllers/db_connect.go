package controllers

import (
	"database/sql"
	"log"
)

var db_connection *sql.DB

func init() {
	var err error
	db_connection, err = sql.Open("sqlite3", "file:./database/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to databse sucessfully")
}
