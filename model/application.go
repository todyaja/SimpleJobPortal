package model

type Application struct {
	Id          int
	JobId       int
	ApplicantId int
	Status      int
}

type ApplicationDetail struct {
	Id       int
	JobId    int
	JobTitle string
	Status   string
}

type JobApplicant struct {
	JobId      int
	JobTitle   string
	Applicants []Applicant
}

type Applicant struct {
	UserId int
	Email  string
	Status string
}
