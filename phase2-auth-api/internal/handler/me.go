package handler

import (
	"net/http"

	"github.com/amirabaris/go-fun/phase2-auth-api/internal/middleware"
	"github.com/amirabaris/go-fun/phase2-auth-api/internal/store"
)

// TODO(phase2-auth-me): Implement protected /me endpoint.
//
// Task:
//   1. MeHandler reads user ID from request context (set by auth middleware).
//   2. GET /me — fetch user by ID, return UserResponse JSON.
//   3. Return 401 if no user in context (shouldn't happen if middleware is correct).
//
// Learn: context values for request-scoped data.

type MeHandler struct {
	store store.Store
}

func NewMeHandler(s store.Store) *MeHandler {
	return &MeHandler{store: s}
}

func (h *MeHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "un authed")
		return
	}

	user, err := h.store.GetUserByID(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch user")
		return
	}
	if user.ID == "" {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	writeJSON(w, http.StatusOK, toUserResponse(user))

}
