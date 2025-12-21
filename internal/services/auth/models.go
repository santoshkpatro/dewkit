package auth

import "time"

type User struct {
	ID                int       `json:"id"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	IsActive          bool      `json:"isActive"`
	PasswordHash      string    `json:"-"`
	Salt              string    `json:"-"`
	IsPasswordExpired bool      `json:"-"`
	Role              string    `json:"role"`
	LastLoginAt       time.Time `json:"lastLoginAt"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}
