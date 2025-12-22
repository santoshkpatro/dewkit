package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoggedInUserResponse struct {
	Email    string `json:"email" db:"email"`
	FullName string `json:"fullName" db:"full_name"`
	Role     string `json:"role" db:"role"`
}
