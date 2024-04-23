package request

type TalentWithdrawApplicationRequest struct {
	ApplicationId int `validate:"required min=1" json:"application_id"`
}
