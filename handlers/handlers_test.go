package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetCandidatesWithAcceptedOffers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(1, "John Doe", "john@example.com", "ABC Company", "1234567890", "offer_rolledout_accepted").
		AddRow(2, "Jane Doe", "jane@example.com", "XYZ Company", "9876543210", "offer_rolledout_accepted")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/candidates_offers_rolledout_accepted", func(ctx *gin.Context) { GetCandidatesWithAcceptedOffers(db, ctx) })

	req, _ := http.NewRequest("GET", "/candidates_offers_rolledout_accepted", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}
func TestGetCandidatesWithAwaitedOffers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(1, "John Doe", "john@example.com", "ABC Company", "1234567890", "offer_rolledout_awaited").
		AddRow(2, "Jane Doe", "jane@example.com", "XYZ Company", "9876543210", "offer_rolledout_awaited")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/candidates_offers_rolledout_awaited", func(ctx *gin.Context) { GetCandidatesWithAwaitedOffers(db, ctx) })

	req, _ := http.NewRequest("GET", "/candidates_offers_rolledout_awaited", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetAcceptedCandidatesCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(2)

	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/count_candidates_offers_rolledout_accepted", func(ctx *gin.Context) { GetAcceptedCandidatesCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/count_candidates_offers_rolledout_accepted", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"accepted_candidates_count\":2}", w.Body.String())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetAwaitedCandidatesCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)

	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/count_candidates_offers_rolledout_awaited", func(ctx *gin.Context) { GetAwaitedCandidatesCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/count_candidates_offers_rolledout_awaited", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"awaited_candidates_count\":1}", w.Body.String())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestWriteToExcel(t *testing.T) {
	// Test data
	filename := "test_output.xlsx"
	candidates := []Candidate{
		{1, "John Doe", "john@example.com", "ABC Company", "1234567890", "offer_rolledout_accepted"},
		{2, "Jane Doe", "jane@example.com", "XYZ Company", "9876543210", "offer_rolledout_accepted"},
	}

	// Execute the writeToExcel function
	err := writeToExcel(filename, candidates)
	defer func() {
		// Clean up: Remove the test output file after the test
		if err := os.Remove(filename); err != nil {
			t.Errorf("Error removing test output file: %v", err)
		}
	}()

	// Assertions
	assert.NoError(t, err, "writeToExcel should not return an error")

	// Additional assertions can be added based on your specific requirements
	// For example, you can check if the file exists, if the headers and data are correctly written, etc.
	// Here, we check if the file exists after writing
	_, err = os.Stat(filename)
	assert.False(t, os.IsNotExist(err), "Test output file should exist")
}
