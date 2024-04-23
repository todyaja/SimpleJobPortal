package request

type UserLoginRequest struct {
	Email    string `validate:"required min=1, max=100" json:"email"`
	Password string `validate:"required min=6, max=100" json:"password"`
}
