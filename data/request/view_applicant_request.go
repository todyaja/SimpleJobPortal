package request

type ViewApplicantRequest struct {
	JobId int `validate:"required min=1" json:"job_id"`
}
