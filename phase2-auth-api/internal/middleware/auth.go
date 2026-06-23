package middleware

import (
	"context"
	"net/http"
)

// TODO(phase2-auth-middleware): Implement authentication middleware.
//
// CHOOSE ONE strategy and implement fully before attempting the other:
//
// Option A — Session cookie (recommended first):
//   1. On login, set HttpOnly, Secure (in prod), SameSite=Lax cookie with session ID.
//   2. Auth middleware reads cookie, looks up session in store, sets user ID in context.
//   3. Use an unexported context key type to avoid collisions.
//
// Option B — Bearer JWT:
//   1. On login, sign JWT with HMAC-SHA256 using secret from config.
//   2. Auth middleware reads Authorization: Bearer <token>, validates, sets user ID in context.
//   3. Use github.com/golang-jwt/jwt is NOT allowed in Phase 2 — implement minimal JWT
//      parsing with crypto/hmac + encoding/json, or stick with sessions.
//
// Task:
//   1. Define context key and helper: UserIDFromContext(ctx) (string, bool).
//   2. Implement Auth(store) func(http.Handler) http.Handler.
//   3. Return 401 JSON if unauthenticated.
//
// Learn: https://pkg.go.dev/context — WithValue, custom key types.

type contextKey string

const userIDKey contextKey = "userID"

func UserIDFromContext(ctx context.Context) (string, bool) {
	panic("TODO: implement UserIDFromContext")
}

func Auth(next http.Handler) http.Handler {
	panic("TODO: implement Auth middleware — choose session OR JWT")
}
