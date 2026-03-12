package request

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type VerifyEmailRequest struct {
	VerificationToken string `json:"verification_token" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type GoogleExchange struct {
	Code  string `json:"code" validate:"required"`
	State string `json:"state" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token           string `json:"token" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UpdateUserProfileRequest struct {
	Name            string `json:"name" validate:"omitempty,min=3"`
	Password        string `json:"password" validate:"omitempty,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"omitempty,eqfield=Password"`
}
