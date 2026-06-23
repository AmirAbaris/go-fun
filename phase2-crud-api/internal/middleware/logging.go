package middleware

import "net/http"

// TODO(phase2-crud-middleware): Implement logging and recover middleware.
//
// Task: Same as phase2-todo-api — Logging, Recover, Chain.

func Logging(next http.Handler) http.Handler {
	panic("TODO: implement Logging middleware")
}

func Recover(next http.Handler) http.Handler {
	panic("TODO: implement Recover middleware")
}
