package store

import (
	"context"

	"github.com/amirabaris/go-fun/phase2-auth-api/internal/domain"
)

// TODO(phase2-auth-store): Define Store interface and in-memory implementation.
//
// Task:
//   1. UserStore methods:
//        CreateUser(ctx, input domain.RegisterInput, passwordHash string) (domain.User, error)
//        GetUserByEmail(ctx, email string) (domain.User, error)
//        GetUserByID(ctx, id string) (domain.User, error)
//   2. SessionStore methods (if using session cookies):
//        CreateSession(ctx, userID string) (sessionID string, err error)
//        GetSessionUserID(ctx, sessionID string) (userID string, err error)
//        DeleteSession(ctx, sessionID string) error
//   3. Return ErrEmailTaken, ErrInvalidCredentials, ErrNotFound as sentinel errors.
//   4. Use sync.RWMutex for concurrent safety.
//
// If using JWT instead of sessions, you may omit SessionStore methods.

type Store interface {
	CreateUser(ctx context.Context, input domain.RegisterInput, passwordHash string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	CreateSession(ctx context.Context, userID string) (string, error)
	GetSessionUserID(ctx context.Context, sessionID string) (string, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

func NewMemoryStore() Store {
	panic("TODO: implement NewMemoryStore")
}
