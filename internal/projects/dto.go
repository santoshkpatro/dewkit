package projects

import "time"

type ProjectCreateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type ProjectListResponse struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	Code        string    `json:"code" db:"code"`
	MemberRole  string    `json:"memberRole" db:"member_role"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type ProjectMemberResponse struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	FullName string `json:"fullName" db:"full_name"`
	Role     string `json:"role" db:"role"`
}
