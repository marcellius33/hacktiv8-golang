package params

type UserCreateRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
