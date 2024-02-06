package models

type Candidate struct {
	CandidateID     int    `json:"candidate_id"`
	Name            string `json:"name"`
	EmailID         string `json:"email_id"`
	CurrentCompany  string `json:"current_company"`
	Mobile          string `json:"mobile"`
	InterviewStatus string `json:"interview_status"`
}

type ResumeCandidate struct {
	CandidateID        int    `json:"candidate_id"`
	Name               string `json:"name"`
	EmailID            string `json:"email_id"`
	CurrentCompany     string `json:"current_company"`
	Mobile             string `json:"mobile"`
	ScreeningStatus    string `json:"screening_status"`
	SkillCategory      string `json:"skill_category"`
	NoticePeriod       string `json:"notice_period"`
	TotalExperience    string `json:"total_experience"`
	Date               string `json:"date"`
	RelevantExperience string `json:"relevent_experience"`
	Comment            string `json:"comment"`
}
