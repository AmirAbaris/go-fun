package domain

import "time"

// TODO(phase3-domain-user): Define User entity.
//
// Task:
//   1. User struct: ID, Email, PasswordHash (json:"-"), CreatedAt.
//   2. RegisterInput, LoginInput with Validate() — reuse patterns from phase2-auth-api.
//   3. UserResponse DTO for API (no password fields).
//   4. Keep package free of infrastructure imports.

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type RegisterInput struct {
	Email    string
	Password string
}

func (in RegisterInput) Validate() error {
	panic("TODO: implement validation")
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
