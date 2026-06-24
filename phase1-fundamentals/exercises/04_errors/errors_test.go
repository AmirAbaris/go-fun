package errors

import (
	"errors"
	"testing"
)

// TODO(phase1-04-test): Test error wrapping and sentinel errors.
//
// Task:
//   1. Test FindUser returns ErrNotFound for invalid id.
//   2. Test GetUser wraps ErrNotFound — use errors.Is to assert.
//   3. Use table-driven style.
//
// Verify: go test ./exercises/04_errors -v

// var ErrNotFound = errors.New("not found")

func TestFindUserNotFound(t *testing.T) {
	_, err := FindUser(0)

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestGetUserWrapsErrNotFound(t *testing.T) {
	_, err := GetUser(0)

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected wrapped ErrNotFound, got %v", err)
	}
}
