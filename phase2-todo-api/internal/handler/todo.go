package handler

import (
	"net/http"

	"github.com/amirabaris/go-fun/phase2-todo-api/internal/store"
)

// TODO(phase2-todo-handler): Implement HTTP handlers for Todo CRUD.
//
// Task:
//   1. Define TodoHandler struct holding store store.Store (injected via constructor).
//   2. Implement NewTodoHandler(s store.Store) *TodoHandler.
//   3. Implement handlers:
//        ListTodos  — GET  /todos       → JSON array
//        CreateTodo — POST /todos       → decode JSON body, validate, 201
//        GetTodo    — GET  /todos/{id}  → 404 if not found
//        UpdateTodo — PATCH /todos/{id} → partial update
//        DeleteTodo — DELETE /todos/{id}→ 204
//   4. Write helper writeJSON(w, status, v) and writeError(w, status, msg).
//   5. Parse id from r.PathValue("id") (Go 1.22+) or manual path parsing.
//
// Learn: encoding/json, net/http status codes, dependency injection.
// DI pattern: handler.NewTodoHandler(store) — no package-level store variable.

type TodoHandler struct {
	store store.Store
}

func NewTodoHandler(s store.Store) *TodoHandler {
	panic("TODO: implement NewTodoHandler")
}

func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement ListTodos")
}
