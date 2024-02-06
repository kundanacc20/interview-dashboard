package models

type InterviewAtStatus struct {
	CandidateID     int    `json:"candidate_id"`
	InterviewStatus string `json:"interview_status"`
	L1ScheduledDate string `json:"L1_scheduled_date"`
	L1Panel         string `json:"L1_panel"`
	L2ScheduledDate string `json:"L2_scheduled_date"`
	L2Panel         string `json:"L2_Panel"`
	DMScheduledDate string `json:"DM_scheduled_date"`
	DMPanel         string `json:"DM_panel"`
	OnboardingDate  string `json:"onboarding_date"`
	Comment         string `json:"comment"`
}
