package middleware

import "net/http"

// TODO(phase2-todo-middleware): Implement HTTP middleware.
//
// Task:
//   1. Logging — log method, path, status code, duration after next.ServeHTTP.
//   2. Recover — defer/recover panic, return 500 JSON error.
//   3. RequestID — generate ID (crypto/rand + hex), set X-Request-ID header.
//   4. Chain helper: func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler
//
// Learn: middleware pattern — func(http.Handler) http.Handler
// Usage: middleware.Chain(handler, middleware.Logging, middleware.Recover)

func Logging(next http.Handler) http.Handler {
	panic("TODO: implement Logging middleware")
}

func Recover(next http.Handler) http.Handler {
	panic("TODO: implement Recover middleware")
}

func RequestID(next http.Handler) http.Handler {
	panic("TODO: implement RequestID middleware")
}

func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	panic("TODO: implement Chain")
}
