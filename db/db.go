package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB
var err error

// databse connection
func InitDB(dsn string) *sql.DB {
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}

	return Db
}
