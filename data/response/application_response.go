package response

type ApplicationDetailResponse struct {
	Id       int    `json:"id"`
	JobId    int    `json:"job_id"`
	JobTitle string `json:"job_title"`
	Status   string `json:"status"`
}

type JobApplicantResponse struct {
	JobId      int         `json:"job_id"`
	JobTitle   string      `json:"job_title"`
	Applicants []Applicant `json:"applicants"`
}

type Applicant struct {
	UserId int    `json:"job_id"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
