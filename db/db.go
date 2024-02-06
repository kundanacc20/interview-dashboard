// db/db.go

package db

import (
	"database/sql"
	"log"
)

// Database represents the database connection
type Database struct {
	Conn *sql.DB
}

// Exec implements handlers.DBHandler.
func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.Conn.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

// InitDB initializes a new Database instance
func InitDB(dsn string) *Database {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Database{Conn: db}
}

// Query executes a SQL query and returns the result
func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.Conn.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rows, nil
}

// QueryRow executes a SQL query that is expected to return at most one row
func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	row := d.Conn.QueryRow(query, args...)
	return row
}
