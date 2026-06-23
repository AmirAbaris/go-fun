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

func TestFindUserNotFound(t *testing.T) {
	t.Skip("implement tests and remove this skip")
	_ = errors.Is
}
