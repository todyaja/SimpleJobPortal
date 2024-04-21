package request

type JobCreateRequest struct {
	Title       string `validate:"required min=1, max=100" json:"title"`
	Detail      string `validate:"required min=1, max=100" json:"detail"`
	Requirement string `validate:"required min=1, max=100" json:"requirement"`
}
