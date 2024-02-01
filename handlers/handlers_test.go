package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	_, err = os.Stat(filename)
	assert.False(t, os.IsNotExist(err), "Test output file should exist")
}

// MockDB is a mock implementation of the DBHandler interface
type MockDB struct {
	mock.Mock
}

// Query is a mock implementation for the Query method in DBHandler
func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	arguments := m.Called(query, args)
	return nil, arguments.Error(1)
}

// QueryRow is a mock implementation for the QueryRow method in DBHandler
func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*sql.Row)
}

func TestGetCandidatesWithAcceptedOffersNegative(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	r := gin.Default()
	r.GET("/candidates_offers_rolledout_accepted", func(ctx *gin.Context) { GetCandidatesWithAcceptedOffers(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/candidates_offers_rolledout_accepted", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}

func TestGetCandidatesWithAwaitedOffersNegative(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	r := gin.Default()
	r.GET("/candidates_offers_rolledout_awaited", func(ctx *gin.Context) { GetCandidatesWithAwaitedOffers(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/candidates_offers_rolledout_awaited", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}
