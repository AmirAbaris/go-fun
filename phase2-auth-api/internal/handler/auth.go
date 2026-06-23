package handler

import (
	"net/http"

	"github.com/amirabaris/go-fun/phase2-auth-api/internal/store"
)

// TODO(phase2-auth-handler): Implement register, login, logout handlers.
//
// Task:
//   1. AuthHandler holds store store.Store (injected).
//   2. Register — POST /register:
//        decode JSON, validate, hash password, create user, return 201 UserResponse.
//   3. Login — POST /login:
//        verify credentials, create session (or issue JWT), set cookie (or return token JSON).
//   4. Logout — POST /logout:
//        delete session, clear cookie (or client discards JWT).
//   5. Return 409 for duplicate email, 401 for bad credentials.
//
// Learn: never return password hash in responses.

type AuthHandler struct {
	store store.Store
}

func NewAuthHandler(s store.Store) *AuthHandler {
	panic("TODO: implement NewAuthHandler")
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement Register")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement Login")
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement Logout")
}
