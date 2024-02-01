// File: db_test.go
package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// TestInitDB tests the InitDB function
func TestInitDB(t *testing.T) {
	// Disable automatic monitoring pings
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(false))
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Set up expectations for the mock
	mock.ExpectPing()

	// Call the function being tested
	InitDB("root:Itsmypasword@047@tcp(localhost:3306)/interview_dashboard1")

	// Assert that expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations: %s", err)
	}
}
