package handler

import (
	"net/http"

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
	panic("TODO: implement NewMeHandler")
}

func (h *MeHandler) Me(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement Me")
}
