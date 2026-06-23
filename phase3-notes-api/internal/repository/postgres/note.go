package postgres

import (
	"context"

	"github.com/amirabaris/go-fun/phase3-notes-api/internal/domain"
	"github.com/amirabaris/go-fun/phase3-notes-api/internal/repository"
)

// TODO(phase3-repo-postgres): Implement NoteRepository using sqlc-generated code.
//
// Milestone 2 — after migrations + sqlc generate.
//
// Task:
//   1. go get github.com/jackc/pgx/v5
//   2. NoteRepository struct holds *pgxpool.Pool and *db.Queries (sqlc generated).
//   3. NewNoteRepository(pool *pgxpool.Pool) repository.NoteRepository.
//   4. Map sqlc row types → domain.Note (don't leak db types to service layer).
//   5. Enforce user_id scoping on every query (users can only access their notes).
//   6. Return domain.ErrNotFound when appropriate.
//   7. Write integration test with local Postgres or testcontainers.
//
// Learn: https://github.com/jackc/pgx, https://docs.sqlc.dev

type NoteRepository struct {
	// pool    *pgxpool.Pool   // add after: go get github.com/jackc/pgx/v5
	// queries *db.Queries     // add after: sqlc generate
}

func NewNoteRepository(pool interface{}) repository.NoteRepository {
	panic("TODO: implement NewNoteRepository(pool *pgxpool.Pool) with sqlc Queries")
}

func (r *NoteRepository) Create(ctx context.Context, userID string, input domain.CreateNoteInput) (domain.Note, error) {
	panic("TODO: implement Create")
}
