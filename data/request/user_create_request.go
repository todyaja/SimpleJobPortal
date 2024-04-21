package request

type UserCreateRequest struct {
	Username string `validate:"required min=1, max=100" json:"username"`
	Email    string `validate:"required min=1, max=100" json:"email"`
	Password string `validate:"required min=6, max=100" json:"password"`
	UserType int    `validate:"required min=1" json:"user_type"`
}
