package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// var Db *sql.DB
// var err error

// databse connection
func InitDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
