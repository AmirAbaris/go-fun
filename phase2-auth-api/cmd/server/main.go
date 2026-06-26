package main

import (
	"log"
	"net/http"

	"github.com/amirabaris/go-fun/phase2-auth-api/internal/config"
	"github.com/amirabaris/go-fun/phase2-auth-api/internal/handler"
	"github.com/amirabaris/go-fun/phase2-auth-api/internal/middleware"
	"github.com/amirabaris/go-fun/phase2-auth-api/internal/store"
)

// TODO(phase2-auth-main): Wire dependencies and start the HTTP server.
//
// Task:
//   1. Load config from internal/config.
//   2. Create in-memory store: store.NewMemoryStore().
//   3. Create handlers: auth.NewAuthHandler(store), handler.NewMeHandler(store).
//   4. Register public routes: POST /register, POST /login.
//   5. Register protected routes behind middleware.Auth:
//        GET  /me
//        POST /logout
//   6. Add logging + recover middleware (reuse pattern from phase2-todo-api).
//   7. Graceful shutdown on SIGINT/SIGTERM.
//
// Run: go run ./cmd/server

func main() {
	s := store.NewMemoryStore()
	authHandler := handler.NewAuthHandler(s)
	meHandler := handler.NewMeHandler(s)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	protected := middleware.Auth(s)
	mux.Handle("GET /me", protected(http.HandlerFunc(meHandler.Me)))
	mux.Handle("POST /logout", protected(http.HandlerFunc(authHandler.Logout)))

	_, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

}
