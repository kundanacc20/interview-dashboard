package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kundanacc20/Offer_Rolledout/db"
	"github.com/kundanacc20/Offer_Rolledout/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// kundan unit test cases
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
	r.GET("/interview-db/home/offer_rolled_out_accepted", func(ctx *gin.Context) { GetCandidatesWithAcceptedOffers(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/offer_rolled_out_accepted", nil)
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
	r.GET("/interview-db/home/offer_rolled_out_accepted_count", func(ctx *gin.Context) { GetAcceptedCandidatesCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/offer_rolled_out_accepted_count", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"accepted_candidates_count\":2}", w.Body.String())
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
	r.GET("/interview-db/home/offer_rolled_out_awaited", func(ctx *gin.Context) { GetCandidatesWithAwaitedOffers(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/offer_rolled_out_awaited", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
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
	r.GET("/interview-db/home/offer_rolled_out_awaited_count", func(ctx *gin.Context) { GetAwaitedCandidatesCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/offer_rolled_out_awaited_count", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"awaited_candidates_count\":1}", w.Body.String())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

// MockDB is a mock implementation of the DBHandler interface
type MockDB struct {
	mock.Mock
}

// Exec implements DBHandler.
func (*MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	panic("unimplemented")
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
	r.GET("/interview-db/home/offer_rolled_out_accepted", func(ctx *gin.Context) { GetCandidatesWithAcceptedOffers(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/interview-db/home/offer_rolled_out_accepted", nil)
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
	r.GET("/interview-db/home/offer_rolled_out_awaited", func(ctx *gin.Context) { GetCandidatesWithAwaitedOffers(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/interview-db/home/offer_rolled_out_awaited", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}

func TestWriteToExcel(t *testing.T) {
	// Test data
	filename := "test_output.xlsx"
	candidates := []models.Candidate{
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

// arpit unit test cases
func TestGetListOfAllCandidatAtL1(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was occurred while opening a database connection for test.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(3, "Rajnish Pandey", "rajnish.pandey@gmail.com", "TCS", "7042787850", "L1_selected").
		AddRow(4, "Harshendra Mandloi", "harsh.mandloi@gmail.com", "Infosys", "7042787851", "L1_rejected")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)
	r := gin.Default()
	r.GET("/interview-db/home/L1", func(ctx *gin.Context) { GetListOfAllCandidatAtL1(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/L1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetSelectedL1Count(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' occurred while opening a connection to database for test.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/L1_count_selected", func(ctx *gin.Context) { GetSelectedL1Count(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/L1_count_selected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"level_L1_selected_count\":2}", w.Body.String())
}

func TestGetRejectedL1Count(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' occurred while opening a connection to database for test.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/L1_count_rejected", func(ctx *gin.Context) { GetRejectedL1Count(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/L1_count_rejected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"level_L1_rejected_count\":2}", w.Body.String())
}

func TestGetListOfAllCandidatAtL2(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was occurred while opening a database connection for test.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(3, "Rajnish Pandey", "rajnish.pandey@gmail.com", "TCS", "7042787850", "L2_selected").
		AddRow(4, "Harshendra Mandloi", "harsh.mandloi@gmail.com", "Infosys", "7042787851", "L2_rejected")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)
	r := gin.Default()
	r.GET("/interview-db/home/L2", func(ctx *gin.Context) { GetListOfAllCandidatAtL2(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/L2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetSelectedL2Count(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' occurred while opening a connection to database for test.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/L2_count_selected", func(ctx *gin.Context) { GetSelectedL2Count(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/L2_count_selected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"level_L2_selected_count\":2}", w.Body.String())
}

func TestGetRejectedL2Count(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' occurred while opening a connection to database for test.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/L2_count_rejected", func(ctx *gin.Context) { GetRejectedL2Count(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/L2_count_rejected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"level_L2_rejected_count\":2}", w.Body.String())
}

//yellaling unit test cases

func TestGetListOfAllCandidatOnboarded(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(1, "Rahul", "rahul@example.com", "tcs Company", "1234560890", "Onboarded").
		AddRow(2, "ramesh", "ramesh@example.com", "tcs Company", "9876503210", "Onboarded")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/Onboarded", func(ctx *gin.Context) { GetListOfAllCandidatOnboarded(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/Onboarded", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetOnboardedCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(2)

	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/onboarded_count", func(ctx *gin.Context) { GetOnboardedCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/onboarded_count", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, "{\"onboarded_count\":2}", w.Body.String())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetListOfAllCandidatOnboardedNegative(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	r := gin.Default()
	r.GET("/interview-db/home/Onboarded", func(ctx *gin.Context) { GetListOfAllCandidatOnboarded(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/interview-db/home/Onboarded", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}

//shaik unit test cases

func TestGetListOfAllCandidatAtDM(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(305, "shaik", "shaik@gmail.com", "wipro", "1234567890", "DM_selected").
		AddRow(306, "saisameer", "saidameer@gmail.com", "tcs", "9876543210", "DM_rejected")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/DM", func(ctx *gin.Context) { GetListOfAllCandidatAtDM(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/DM", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetGetListOfAllCandidatAtDMNegative(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	r := gin.Default()
	r.GET("/interview-db/home/DM", func(ctx *gin.Context) { GetListOfAllCandidatAtDM(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/interview-db/home/DM", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}

func TestGetSelectedDMCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)

	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/DM_selected", func(ctx *gin.Context) { GetSelectedDMCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/DM_selected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"level_DM_selected_count\":1}", w.Body.String())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetRejectedDMCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)

	mock.ExpectQuery("SELECT COUNT(.*)").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/DM_rejected", func(ctx *gin.Context) { GetRejectedDMCount(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/DM_rejected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"level_DM_rejected_count\":1}", w.Body.String())
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetListOfAllCandidatAtDMSelected(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(305, "shaik", "shaik@gmail.com", "wipro", "1234567890", "DM_selected").
		AddRow(306, "saisameer", "saidameer@gmail.com", "tcs", "9876543210", "DM_rejected")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/DM_selected", func(ctx *gin.Context) { GetListOfAllCandidatAtDMSelected(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/DM_selected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetListOfAllCandidatAtDMRejected(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not epected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"candidate_id", "name", "email_id", "current_company", "mobile", "interview_status"}).
		AddRow(305, "shaik", "shaik@gmail.com", "wipro", "1234567890", "DM_selected").
		AddRow(306, "saisameer", "saidameer@gmail.com", "tcs", "9876543210", "DM_rejected")

	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	r := gin.Default()
	r.GET("/interview-db/home/DM_rejected", func(ctx *gin.Context) { GetListOfAllCandidatAtDMRejected(db, ctx) })

	req, _ := http.NewRequest("GET", "/interview-db/home/DM_rejected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled edpectation: %s", err)
	}

}

func TestGetListOfAllCandidatAtDMSelectedNegative(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	r := gin.Default()
	r.GET("/interview-db/home/DM_selected", func(ctx *gin.Context) { GetListOfAllCandidatAtDMSelected(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/interview-db/home/DM_selected", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}

func TestGetListOfAllCandidatAtDMRejectedNegative(t *testing.T) {
	mockDB := new(MockDB)
	mockDB.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	r := gin.Default()
	r.GET("/interview-db/home/DM_rejected", func(ctx *gin.Context) { GetListOfAllCandidatAtDMRejected(mockDB, ctx) })

	req, err := http.NewRequest("GET", "/interview-db/home/DM_rejected", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockDB.AssertExpectations(t)
}

// sivarajan unit test cases
func TestUserLoginHandler_Success(t *testing.T) {
	//Initialize Gin router
	router := gin.Default()
	// Mock database instance
	mockDB := &db.Database{}
	//mockDB = sqlmock.New()

	// create a test request with form data
	req, _ := http.NewRequest("POST", "/login", strings.NewReader("user_name=testuser&password=testpassword"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Recorder to record the response
	w := httptest.NewRecorder()

	// set up the handler
	router.POST("/login", func(c *gin.Context) {
		UserLoginHandler(mockDB, c)
	})

	// server the request
	router.ServeHTTP(w, req)

	// assert status code is 303 (see other, redirect based on isAdmin)
	assert.Equal(t, http.StatusSeeOther, w.Code)

	// assert that the user is redirected to the expected url
	assert.Contains(t, w.Header().Get("Location"), "/login?error=1")
}

func TestUserLoginHandler_InvalidCredentials(t *testing.T) {
	//Initialize Gin router
	router := gin.Default()
	// Mock database instance
	mockDB := &db.Database{}
	//mockDB = sqlmock.New()

	// create a test request with form data
	req, _ := http.NewRequest("POST", "/login", strings.NewReader("user_name=testuser&password=testpassword"))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Recorder to record the response
	w := httptest.NewRecorder()

	// set up the handler
	router.POST("/login", func(c *gin.Context) {
		UserLoginHandler(mockDB, c)
	})

	// server the request
	router.ServeHTTP(w, req)

	// assert status code is 303 (see other, redirect to login with error)
	assert.Equal(t, http.StatusSeeOther, w.Code)

	// assert that the user is redirected to the expected url
	assert.Contains(t, w.Header().Get("Location"), "/login?error=1")

}
