package middleware

import (
	"context"
	"net/http"
)

// TODO(phase3-middleware): Implement auth and logging middleware.
//
// Milestone 4 — reuse patterns from phase2-auth-api.
//
// Task:
//   1. Auth middleware — session cookie or JWT (same choice as Phase 2).
//   2. UserIDFromContext helper with unexported context key type.
//   3. Logging + Recover middleware.
//   4. Chain helper for composing middleware.
//
// Learn: middleware runs before handler; auth sets user ID in context.

type contextKey string

const userIDKey contextKey = "userID"

func UserIDFromContext(ctx context.Context) (string, bool) {
	panic("TODO: implement UserIDFromContext")
}

func Auth(next http.Handler) http.Handler {
	panic("TODO: implement Auth middleware")
}

func Logging(next http.Handler) http.Handler {
	panic("TODO: implement Logging middleware")
}

func Recover(next http.Handler) http.Handler {
	panic("TODO: implement Recover middleware")
}
