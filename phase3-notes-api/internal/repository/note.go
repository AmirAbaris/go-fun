package repository

import (
	"context"

	"github.com/amirabaris/go-fun/phase3-notes-api/internal/domain"
)

// TODO(phase3-repo-interface): Define repository interfaces (ports).
//
// Milestone 2 — service layer depends on THESE interfaces, not postgres types.
//
// Task:
//   1. NoteRepository interface:
//        Create(ctx, userID string, input domain.CreateNoteInput) (domain.Note, error)
//        GetByID(ctx, userID, noteID string) (domain.Note, error)
//        ListByUser(ctx, userID string) ([]domain.Note, error)
//        ListByUserAndTag(ctx, userID, tag string) ([]domain.Note, error)
//        Update(ctx, userID, noteID string, input domain.UpdateNoteInput) (domain.Note, error)
//        Delete(ctx, userID, noteID string) error
//   2. UserRepository interface (for auth):
//        CreateUser, GetByEmail, GetByID
//   3. Implementation lives in repository/postgres/ — NOT here.
//
// Learn: repository pattern — interface in parent package, impl in subpackage.

type NoteRepository interface {
	Create(ctx context.Context, userID string, input domain.CreateNoteInput) (domain.Note, error)
	GetByID(ctx context.Context, userID, noteID string) (domain.Note, error)
	ListByUser(ctx context.Context, userID string) ([]domain.Note, error)
	ListByUserAndTag(ctx context.Context, userID, tag string) ([]domain.Note, error)
	Update(ctx context.Context, userID, noteID string, input domain.UpdateNoteInput) (domain.Note, error)
	Delete(ctx context.Context, userID, noteID string) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, input domain.RegisterInput, passwordHash string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetByID(ctx context.Context, id string) (domain.User, error)
}
