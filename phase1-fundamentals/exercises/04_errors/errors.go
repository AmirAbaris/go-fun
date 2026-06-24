package errors

import (
	"errors"
	"fmt"
)

// TODO(phase1-04): Learn error handling in Go.
//
// Task:
//   1. Define sentinel error var ErrNotFound = errors.New("not found").
//   2. Write FindUser(id int) (*User, error) that returns ErrNotFound for id <= 0.
//   3. Write GetUser(id int) (*User, error) that wraps ErrNotFound with fmt.Errorf("...: %w", err).
//   4. Demonstrate errors.Is and errors.As in a caller.
//
// Learn: https://go.dev/blog/error-handling-and-go, https://pkg.go.dev/errors
// Verify: go test ./exercises/04_errors -v

var ErrNotFound = errors.New("not found")

type User struct {
	ID   int
	Name string
}

// cuz *User and error are types, they dont names assigned to themsselves.
func FindUser(id int) (*User, error) {
	if id <= 0 {
		return nil, ErrNotFound
	}

	return &User{
		ID:   id,
		Name: "Amir",
	}, nil
}

func GetUser(id int) (*User, error) {
	user, err := FindUser(id)

	// %w is a format verb
	// fmt.Printf("%d", 42)       // 42
	// fmt.Printf("%s", "hello")  // hello
	// fmt.Printf("%t", true)     // true
	// %w -> wrap inside a new error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
