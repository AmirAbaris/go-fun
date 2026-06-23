package main

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
	panic("TODO: implement server wiring")
}
