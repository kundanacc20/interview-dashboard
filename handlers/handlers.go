package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/kundanacc20/Offer_Rolledout/db"
)

type Candidate struct {
	CandidateID     int    `json:"candidate_id"`
	Name            string `json:"name"`
	EmailID         string `json:"email_id"`
	CurrentCompany  string `json:"current_company"`
	Mobile          string `json:"mobile"`
	InterviewStatus string `json:"interview_status"`
}

// type DBHandler interface {
// 	Query(query string, args ...interface{}) (*sql.Rows, error)
// }

func GetCandidatesWithAcceptedOffers(c *gin.Context) {
	// Query the database for candidates with interview status "offer_rolledout_accepted"
	rows, err := db.Db.Query("SELECT r.candidate_id, r.name, r.email_id, r.current_company, r.mobile, ist.interview_status " +
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
	var candidates []Candidate
	for rows.Next() {
		var candidate Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.Name, &candidate.EmailID, &candidate.CurrentCompany, &candidate.Mobile, &candidate.InterviewStatus); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		candidates = append(candidates, candidate)
	}

	// Return the list of candidates with interview status =offer_rolledout_accepted
	c.JSON(http.StatusOK, candidates)

	// writing data to file
	err = writeToExcel("accepted_candidates.xlsx", candidates)
	if err != nil {
		log.Println("Error writing data to Excel:", err)
	}
}

// function that will get database form database and write data in excel
func writeToExcel(filename string, candidates []Candidate) error {
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
