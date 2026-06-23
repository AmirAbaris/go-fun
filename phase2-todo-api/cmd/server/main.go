package main

// TODO(phase2-todo-main): Wire dependencies and start the HTTP server.
//
// Task:
//   1. Load config from internal/config.
//   2. Create in-memory store: store.NewMemoryStore().
//   3. Create handler: handler.NewTodoHandler(store).
//   4. Build middleware chain: logging → recover → handler.
//   5. Register routes on http.NewServeMux:
//        GET  /health
//        GET  /todos
//        POST /todos
//        GET  /todos/{id}
//        PATCH /todos/{id}
//        DELETE /todos/{id}
//   6. Start http.Server with ReadHeaderTimeout.
//   7. Implement graceful shutdown on SIGINT/SIGTERM (context.WithTimeout + server.Shutdown).
//
// Learn: https://pkg.go.dev/net/http#Server
// Run: go run ./cmd/server

func main() {
	panic("TODO: implement server wiring")
}
