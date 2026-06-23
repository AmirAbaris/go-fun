package store

import (
	"context"

	"github.com/amirabaris/go-fun/phase2-todo-api/internal/domain"
)

// TODO(phase2-todo-store): Define Store interface and in-memory implementation.
//
// Task:
//   1. Define Store interface with methods:
//        List(ctx context.Context) ([]domain.Todo, error)
//        Get(ctx context.Context, id string) (domain.Todo, error)
//        Create(ctx context.Context, input domain.CreateTodoInput) (domain.Todo, error)
//        Update(ctx context.Context, id string, input domain.UpdateTodoInput) (domain.Todo, error)
//        Delete(ctx context.Context, id string) error
//   2. Implement MemoryStore with sync.RWMutex + map[string]domain.Todo.
//   3. Generate IDs with crypto/rand or google/uuid is NOT allowed — use stdlib only
//      (e.g. encoding/hex + crypto/rand, or incremental IDs for simplicity).
//   4. Return a sentinel error ErrNotFound when id doesn't exist.
//
// Learn: sync.RWMutex, context.Context as first param convention.

type Store interface {
	List(ctx context.Context) ([]domain.Todo, error)
	Get(ctx context.Context, id string) (domain.Todo, error)
	Create(ctx context.Context, input domain.CreateTodoInput) (domain.Todo, error)
	Update(ctx context.Context, id string, input domain.UpdateTodoInput) (domain.Todo, error)
	Delete(ctx context.Context, id string) error
}

func NewMemoryStore() Store {
	panic("TODO: implement NewMemoryStore")
}
