package middleware

import "net/http"

// TODO(phase2-auth-logging): Implement logging and recover middleware.
//
// Task: Same pattern as phase2-todo-api — Logging, Recover, Chain helpers.
// Reuse what you built in the Todo API or copy the approach here.

func Logging(next http.Handler) http.Handler {
	panic("TODO: implement Logging middleware")
}

func Recover(next http.Handler) http.Handler {
	panic("TODO: implement Recover middleware")
}
