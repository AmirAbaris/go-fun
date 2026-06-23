package store

import (
	"context"
	"database/sql"

	"github.com/amirabaris/go-fun/phase2-crud-api/internal/domain"
)

// TODO(phase2-crud-store): Implement PostgreSQL store with database/sql.
//
// Task:
//   1. PostgresStore holds *sql.DB (injected).
//   2. Implement Store interface:
//        List(ctx) ([]domain.Product, error)
//        Get(ctx, id int64) (domain.Product, error)
//        Create(ctx, input) (domain.Product, error)
//        Update(ctx, id, input) (domain.Product, error)
//        Delete(ctx, id int64) error
//   3. Use QueryContext, ExecContext, QueryRowContext — NEVER fmt.Sprintf for SQL values.
//   4. Use $1, $2 placeholders for all parameters.
//   5. Return ErrNotFound when sql.ErrNoRows.
//   6. Scan rows into domain.Product.
//
// Learn: https://pkg.go.dev/database/sql, https://go.dev/doc/database/querying

type Store interface {
	List(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int64) (domain.Product, error)
	Create(ctx context.Context, input domain.CreateProductInput) (domain.Product, error)
	Update(ctx context.Context, id int64, input domain.UpdateProductInput) (domain.Product, error)
	Delete(ctx context.Context, id int64) error
}

func NewPostgresStore(db *sql.DB) Store {
	panic("TODO: implement NewPostgresStore")
}
