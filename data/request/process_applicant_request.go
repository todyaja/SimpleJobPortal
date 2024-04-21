package request

type ProcessApplicantRequest struct {
	AppplicationId int `validate:"required min=1, max=100" json:"application_id"`
	ChangeStatusTo int `validate:"required min=1, max=100" json:"change_status_to"`
}
