package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amirabaris/go-fun/phase2-auth-api/internal/store"
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
const sessionCookieName = "session_id"

func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)

	return userID, ok
}

func Auth(s store.Store) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(sessionCookieName)
			if err != nil {
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}

			userID, err := s.GetSessionUserID(r.Context(), cookie.Value)
			if err != nil || userID == "" {
				writeUnauthorized(w)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userID)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func writeUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
}
