package handler

import (
	"net/http"

	"github.com/amirabaris/go-fun/phase2-crud-api/internal/store"
)

// TODO(phase2-crud-handler): Implement Product HTTP handlers.
//
// Task:
//   1. ProductHandler with injected store.Store.
//   2. CRUD handlers mirroring phase2-todo-api patterns.
//   3. Parse id as int64 from path — return 400 on invalid id.
//   4. Consistent JSON error format: {"error": "message"}.
//
// Reuse patterns from your Todo API implementation.

type ProductHandler struct {
	store store.Store
}

func NewProductHandler(s store.Store) *ProductHandler {
	panic("TODO: implement NewProductHandler")
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement ListProducts")
}
