package domain

import (
	"errors"
	"time"
)

// TODO(phase2-auth-domain): Define User domain model.
//
// Task:
//   1. Define User struct: ID, Email, CreatedAt. Never expose PasswordHash in JSON responses.
//   2. Define RegisterInput: Email, Password with Validate() — email non-empty, password min 8 chars.
//   3. Define LoginInput: Email, Password.
//   4. Define UserResponse DTO for API responses (no password fields).
//   5. Keep this package free of net/http and bcrypt imports.
//
// Learn: separate domain types from HTTP/DB concerns.

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
	if in.Email == "" {
		return errors.New("email is required")
	}
	if len(in.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (in LoginInput) Validate() error {
	if in.Email == "" {
		return errors.New("email is required")
	}
	if len(in.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
