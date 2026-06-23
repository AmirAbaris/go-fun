package errors

import "errors"

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

func FindUser(id int) (*User, error) {
	panic("TODO: implement FindUser")
}

func GetUser(id int) (*User, error) {
	panic("TODO: implement GetUser")
}
