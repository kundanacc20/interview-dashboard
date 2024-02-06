package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kundanacc20/Offer_Rolledout/db"
	"github.com/kundanacc20/Offer_Rolledout/models"
)

type DBHandler interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

//kundan task

func GetCandidatesWithAcceptedOffers(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
		"FROM resume r " +
		"JOIN interview_status_table ist ON r.candidate_id = ist.candidate_id " +
		"WHERE ist.interview_status = 'offer_rolledout_accepted'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.Candidate
	for rows.Next() {
		var candidate models.Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =offer_rolledout_accepted
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcel("accepted_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}
}

func GetCandidatesWithAwaitedOffers(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_awaited"
	rows, err := db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
		"FROM resume r " +
		"JOIN interview_status_table ist ON r.candidate_id = ist.candidate_id " +
		"WHERE ist.interview_status = 'offer_rolledout_awaited'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.Candidate
	for rows.Next() {
		var candidate models.Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status "offer_rolledout_awaited"
	c.JSON(http.StatusOK, candidates)

	// write data to file
	err = writeToExcel("awaited_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}
}

func GetAcceptedCandidatesCount(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'offer_rolledout_accepted'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"accepted_candidates_count": count})
}

func GetAwaitedCandidatesCount(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_awaited"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'offer_rolledout_awaited'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_awaited"
	c.JSON(http.StatusOK, gin.H{"awaited_candidates_count": count})
}

// function that will get database form database and write data in excel
func writeToExcel(filename string, candidates []models.Candidate) error {
	f := excelize.NewFile()

	// Create a new sheet in the Excel file
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// Set the headers in the first row
	headers := []string{"Candidate ID", "Name", "Email ID", "Current Company", "Mobile", "Interview Status"}
	for col, header := range headers {
		cell := excelize.ToAlphaString(col+1) + "1"
		f.SetCellValue(sheetName, cell, header)
	}

	// Populate data in subsequent rows
	for row, candidate := range candidates {
		f.SetCellValue(sheetName, excelize.ToAlphaString(1)+strconv.Itoa(row+2), candidate.CandidateID)
		f.SetCellValue(sheetName, excelize.ToAlphaString(2)+strconv.Itoa(row+2), candidate.Name)
		f.SetCellValue(sheetName, excelize.ToAlphaString(3)+strconv.Itoa(row+2), candidate.EmailID)
		f.SetCellValue(sheetName, excelize.ToAlphaString(4)+strconv.Itoa(row+2), candidate.CurrentCompany)
		f.SetCellValue(sheetName, excelize.ToAlphaString(5)+strconv.Itoa(row+2), candidate.Mobile)
		f.SetCellValue(sheetName, excelize.ToAlphaString(6)+strconv.Itoa(row+2), candidate.InterviewStatus)
	}

	// Set the created sheet as active
	f.SetActiveSheet(f.GetSheetIndex(sheetName))

	// Save the Excel file
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}

// yellaling task
func GetListOfAllCandidatOnboarded(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
		"FROM resume r " +
		"JOIN interview_status_table ist ON r.candidate_id = ist.candidate_id " +
		"WHERE ist.interview_status = 'onboarded'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.Candidate
	for rows.Next() {
		var candidate models.Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcel("onboarded_cadidate_list.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

func GetOnboardedCount(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'onboarded'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"onboarded_candidate_count": count})
}

//shaik saisameer task

func GetListOfAllCandidatAtDM(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
		"FROM resume r " +
		"JOIN interview_status_table ist ON r.candidate_id = ist.candidate_id WHERE ist.interview_status = 'DM_selected' or ist.interview_status= 'DM_rejected'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.Candidate
	for rows.Next() {
		var candidate models.Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcel("candidates_at_DM_level.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

func GetSelectedDMCount(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'DM_selected'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"level_DM_selected_count": count})
}

func GetRejectedDMCount(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'DM_rejected'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"level_DM_rejected_count": count})
}

//arpit task

func GetListOfAllCandidatAtL1(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
		"FROM resume r " +
		"JOIN interview_status_table ist ON r.candidate_id = ist.candidate_id WHERE ist.interview_status = 'L1_selected' or ist.interview_status= 'L1_rejected'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.Candidate
	for rows.Next() {
		var candidate models.Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcel("L1_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

func GetSelectedL1Count(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'L1_selected'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"level_L1_selected_count": count})
}

func GetRejectedL1Count(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'L1_rejected'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"level_L1_rejected_count": count})
}

func GetListOfAllCandidatAtL2(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
		"FROM resume r " +
		"JOIN interview_status_table ist ON r.candidate_id = ist.candidate_id WHERE ist.interview_status = 'L2_selected' or ist.interview_status= 'L2_rejected'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.Candidate
	for rows.Next() {
		var candidate models.Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcel("L2_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

func GetSelectedL2Count(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'L2_selected'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"level_L2_selected_count": count})
}

func GetRejectedL2Count(db DBHandler, c *gin.Context) {
	// Query the database for the count of candidates with interview status "offer_rolledout_accepted"
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_status_table WHERE interview_status = 'L2_rejected'").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}

	// Return the total count of candidates with interview status "offer_rolledout_accepted"
	c.JSON(http.StatusOK, gin.H{"level_L2_rejected_count": count})
}

// anushree task
func GetListOfAllCandidatResumeShortListed(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id,r.skill_category, r.name,r.mobile, r.email_id, r.current_company,r.total_experience,r.relevent_experience,r.notice_period, r.comment, r.screening_status " +
		"FROM resume r WHERE r.screening_status= 'shortlisted'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.ResumeCandidate
	for rows.Next() {
		var candidate models.ResumeCandidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.SkillCategory, &candidate.Name, &candidate.Mobile, &candidate.EmailID, &candidate.CurrentCompany, &candidate.TotalExperience, &candidate.RelevantExperience, &candidate.NoticePeriod, &candidate.Comment, &candidate.ScreeningStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcelResume("Resume_shortlisted_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

func GetListOfAllCandidatResumeRejected(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.candidate_id,r.skill_category, r.name, r.mobile, r.email_id, r.current_company,r.total_experience,r.relevent_experience,r.notice_period, r.comment, r.screening_status " +
		"FROM resume r WHERE r.screening_status= 'rejected'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.ResumeCandidate
	for rows.Next() {
		var candidate models.ResumeCandidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.SkillCategory, &candidate.Name, &candidate.Mobile, &candidate.EmailID, &candidate.CurrentCompany, &candidate.TotalExperience, &candidate.RelevantExperience, &candidate.NoticePeriod, &candidate.Comment, &candidate.ScreeningStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcelResume("Resume_rejected_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

// sindhushree task
func AddCandidateToResume(db DBHandler, C *gin.Context) {
	var candidate models.ResumeCandidate

	// Bind JSON request body to the Candidate struct
	if err := C.ShouldBindJSON(&candidate); err != nil {
		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Insert the candidate into the database
	result, err := db.Exec("INSERT INTO interview_dashboard1.resume (date, skill_category,candidate_id,name, mobile, email_id, total_experience, relevent_experience, current_company, notice_period, comment, screening_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		candidate.Date, candidate.SkillCategory, candidate.CandidateID, candidate.Name, candidate.Mobile, candidate.EmailID, candidate.TotalExperience, candidate.RelevantExperience, candidate.CurrentCompany, candidate.NoticePeriod, candidate.Comment, candidate.ScreeningStatus)
	if err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error inserting candidate into the database:", err)
		return
	}

	// Get the ID of the inserted candidate
	candidateID, _ := result.LastInsertId()
	candidate.CandidateID = int(candidateID)

	// Return the inserted candidate
	C.JSON(http.StatusOK, candidate)

}

func GetListOfAllCandidate(db DBHandler, c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Query("SELECT r.date, r.candidate_id,r.skill_category, r.name, r.email_id, r.current_company,r.total_experience,r.relevent_experience,r.notice_period, r.mobile, r.comment,r.screening_status " +
		"FROM resume r")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying the database:", err)
		return
	}
	defer rows.Close()

	// Iterate through the result set and build a list of candidates
	var candidates []models.ResumeCandidate
	for rows.Next() {
		var candidate models.ResumeCandidate
		if err := rows.Scan(&candidate.Date, &candidate.CandidateID, &candidate.SkillCategory, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.TotalExperience, &candidate.RelevantExperience, &candidate.NoticePeriod, &candidate.Mobile, &candidate.Comment, &candidate.ScreeningStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =onboarded
	c.JSON(http.StatusOK, candidates)

	// writing data to excel file
	err = writeToExcelResume("All_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}

}

// func UpdateCandidate(db DBHandler, C *gin.Context) {
// 	var candidate models.ResumeCandidate

// 	// Bind JSON request body to the Candidate struct
// 	if err := C.ShouldBindJSON(&candidate); err != nil {
// 		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
// 		return
// 	}

// 	// Check if candidate_id is provided in the request
// 	if candidate.CandidateID == 0 {
// 		C.JSON(http.StatusBadRequest, gin.H{"error": "Candidate ID is required for update"})
// 		return
// 	}

// 	// Check if the candidate with the provided candidate_id exists
// 	var existingCandidate models.ResumeCandidate
// 	err := db.QueryRow("SELECT * FROM interview_dashboard1.resume WHERE candidate_id = ?", candidate.CandidateID).
// 		Scan(&existingCandidate.Date, &existingCandidate.SkillCategory, &existingCandidate.CandidateID, &existingCandidate.Name, &existingCandidate.Mobile, &existingCandidate.EmailID, &existingCandidate.TotalExperience, &existingCandidate.RelevantExperience, &existingCandidate.CurrentCompany, &existingCandidate.NoticePeriod, &existingCandidate.Comment, &existingCandidate.ScreeningStatus)

// 	if err == sql.ErrNoRows {
// 		C.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
// 		return
// 	} else if err != nil {
// 		C.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		log.Println("Error querying candidate from the database:", err)
// 		return
// 	}

// 	// Update the candidate in the database
// 	_, err = db.Exec("UPDATE interview_dashboard1.resume SET date=?, skill_category=?, name=?, mobile=?, email_id=?, total_experience=?, relevent_experience=?, current_company=?, notice_period=?, comment=?, screening_status=? WHERE candidate_id=?",
// 		candidate.Date, candidate.SkillCategory, candidate.Name, candidate.Mobile, candidate.EmailID, candidate.TotalExperience, candidate.RelevantExperience, candidate.CurrentCompany, candidate.NoticePeriod, candidate.Comment, candidate.ScreeningStatus, candidate.CandidateID)
// 	if err != nil {
// 		C.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		log.Println("Error updating candidate in the database:", err)
// 		return
// 	}

// 	// Return the updated candidate
// 	C.JSON(http.StatusOK, candidate)
// }

func UpdateCandidateAtInterviewStatus(db DBHandler, C *gin.Context) {
	var candidate models.InterviewAtStatus

	// Bind JSON request body to the Candidate struct
	if err := C.ShouldBindJSON(&candidate); err != nil {
		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check if candidate_id is provided in the request
	if candidate.CandidateID == 0 {
		C.JSON(http.StatusBadRequest, gin.H{"error": "Candidate ID is required for update"})
		return
	}

	// Check if the candidate with the provided candidate_id exists
	var existingCandidate models.InterviewAtStatus
	err := db.QueryRow("SELECT * FROM interview_dashboard1.interview_status_table WHERE candidate_id = ?", candidate.CandidateID).
		Scan(&existingCandidate.CandidateID, &existingCandidate.InterviewStatus, &existingCandidate.L1ScheduledDate, &existingCandidate.L1Panel, &existingCandidate.L2ScheduledDate, &existingCandidate.L2Panel, &existingCandidate.DMScheduledDate, &existingCandidate.DMPanel, &existingCandidate.OnboardingDate, &existingCandidate.Comment)

	if err == sql.ErrNoRows {
		C.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	} else if err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error querying candidate from the database:", err)
		return
	}

	// Update the candidate in the database
	_, err = db.Exec("UPDATE interview_dashboard1.interview_status_table SET interview_status=?, L1_scheduled_date=?, L1_panel=?, L2_scheduled_date=?, L2_panel=?, DM_scheduled_date=?, DM_panel=?, onboarding_date=?, comment=? WHERE candidate_id=?",
		candidate.InterviewStatus, candidate.L1ScheduledDate, candidate.L1Panel, candidate.L2ScheduledDate, candidate.L2Panel, candidate.DMScheduledDate, candidate.DMPanel, candidate.OnboardingDate, candidate.Comment, candidate.CandidateID)
	if err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error updating candidate in the database:", err)
		return
	}

	// Return the updated candidate
	C.JSON(http.StatusOK, candidate)
}

// sathiya task

func UserExists(db DBHandler, username string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM interview_dashboard1.users WHERE user_name = ?", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func SignupHandler(db DBHandler, c *gin.Context) {
	var user models.User

	// Bind JSON request body to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check if the user already exists
	exists, err := UserExists(db, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error checking user existence:", err)
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Insert the user into the database
	result, err := db.Exec("INSERT INTO interview_dashboard1.users (id,user_name, password, is_admin) VALUES (?, ?, ?, ?)",
		user.ID, user.Username, user.Password, user.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		log.Println("Error inserting user into the database:", err)
		return
	}

	// Get the ID of the inserted user
	userID, _ := result.LastInsertId()
	user.ID = int(userID)

	// Return the inserted user
	c.JSON(http.StatusOK, user)
}

//thirlakuppam task

// QueryUserByUsername retrieves user info by username
func QueryUserByUsername(dbInstance *db.Database, username string) (models.User, error) {
	var user models.User

	// Adjust the SQL query according to your table structure
	query := "SELECT user_name, password, is_admin FROM interview_dashboard1.users WHERE user_name = ?"
	row := dbInstance.QueryRow(query, username)

	err := row.Scan(&user.Username, &user.Password, &user.IsAdmin)

	return user, err
}

func UserLoginHandler(db *db.Database, c *gin.Context) {
	// Get username and password from the form data
	username := c.PostForm("user_name")
	password := c.PostForm("password")

	// Open Mysql db connection
	dbInstance, err := sql.Open("mysql", "root:Itsmypasword@047@tcp(localhost:3306)/interview_dashboard1")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//defer dbInstance.Close()

	db.Conn = dbInstance

	// Query user by username
	user, err := QueryUserByUsername(db, username)

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?error=1")
		return
	}

	// Validate password
	if user.Password != password {
		c.Redirect(http.StatusSeeOther, "/login?error=1")
		return
	}

	// Redirect based on isAdmin
	if user.IsAdmin {
		c.Redirect(http.StatusSeeOther, "/interview-dashboard/home/admin")
	} else {
		c.Redirect(http.StatusSeeOther, "/interview-dashboard/home")
	}
}

func writeToExcelResume(filename string, candidates []models.ResumeCandidate) error {
	f := excelize.NewFile()

	// Create a new sheet in the Excel file
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// Set the headers in the first row
	headers := []string{"Candidate ID", "Date", "Skill_Category", "Name", "Mobile", "Email ID", "Total_Experience", "Relevent_Experience", "Current Company", "Notice_Period", "Comment", "Screening_Status"}
	for col, header := range headers {
		cell := excelize.ToAlphaString(col+1) + "1"
		f.SetCellValue(sheetName, cell, header)
	}

	// Populate data in subsequent rows
	for row, candidate := range candidates {
		f.SetCellValue(sheetName, excelize.ToAlphaString(1)+strconv.Itoa(row+2), candidate.CandidateID)
		f.SetCellValue(sheetName, excelize.ToAlphaString(2)+strconv.Itoa(row+2), candidate.Date)
		f.SetCellValue(sheetName, excelize.ToAlphaString(3)+strconv.Itoa(row+2), candidate.SkillCategory)
		f.SetCellValue(sheetName, excelize.ToAlphaString(4)+strconv.Itoa(row+2), candidate.Name)
		f.SetCellValue(sheetName, excelize.ToAlphaString(5)+strconv.Itoa(row+2), candidate.Mobile)
		f.SetCellValue(sheetName, excelize.ToAlphaString(6)+strconv.Itoa(row+2), candidate.EmailID)
		f.SetCellValue(sheetName, excelize.ToAlphaString(7)+strconv.Itoa(row+2), candidate.TotalExperience)
		f.SetCellValue(sheetName, excelize.ToAlphaString(8)+strconv.Itoa(row+2), candidate.RelevantExperience)
		f.SetCellValue(sheetName, excelize.ToAlphaString(9)+strconv.Itoa(row+2), candidate.CurrentCompany)
		f.SetCellValue(sheetName, excelize.ToAlphaString(10)+strconv.Itoa(row+2), candidate.NoticePeriod)
		f.SetCellValue(sheetName, excelize.ToAlphaString(11)+strconv.Itoa(row+2), candidate.Comment)
		f.SetCellValue(sheetName, excelize.ToAlphaString(12)+strconv.Itoa(row+2), candidate.ScreeningStatus)

	}

	// Set the created sheet as active
	f.SetActiveSheet(f.GetSheetIndex(sheetName))

	// Save the Excel file
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
