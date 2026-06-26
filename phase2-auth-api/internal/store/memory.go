package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"

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

var (
	ErrEmailTaken         = errors.New("email already exist")
	ErrInvalidCredentials = errors.New("invalid creds")
	ErrNotFound           = errors.New("not found")
)

type Store interface {
	CreateUser(ctx context.Context, input domain.RegisterInput, passwordHash string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	CreateSession(ctx context.Context, userID string) (string, error)
	GetSessionUserID(ctx context.Context, sessionID string) (string, error)
	DeleteSession(ctx context.Context, sessionID string) error
	GetPasswordHashByUserID(ctx context.Context, userID string) (string, error)
}

type MemoryStore struct {
	mu sync.RWMutex

	usersByID   map[string]domain.User
	userByEmail map[string]domain.User

	passwordHashs map[string]string
	session       map[string]string
}

// GetPasswordHashByUserID implements Store.
func (m *MemoryStore) GetPasswordHashByUserID(ctx context.Context, userID string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	m.mu.RLock()
	defer m.mu.Unlock()

	passwordHash, exists := m.passwordHashs[userID]
	if !exists {
		return "", ErrNotFound
	}

	return passwordHash, nil
}

func randomID() string {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

// CreateSession implements Store.
func (m *MemoryStore) CreateSession(ctx context.Context, userID string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	sessionID := randomID()

	m.mu.Lock()
	defer m.mu.Unlock()

	m.session[sessionID] = userID

	return sessionID, nil
}

// CreateUser implements Store.
func (m *MemoryStore) CreateUser(ctx context.Context, input domain.RegisterInput, passwordHash string) (domain.User, error) {
	if err := ctx.Err(); err != nil {
		return domain.User{}, err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.userByEmail[input.Email]; exists {
		return domain.User{}, ErrEmailTaken
	}

	user := domain.User{
		ID:    randomID(),
		Email: input.Email,
	}

	m.usersByID[user.ID] = user
	m.userByEmail[user.Email] = user
	m.passwordHashs[user.ID] = passwordHash

	return user, nil
}

// DeleteSession implements Store.
func (m *MemoryStore) DeleteSession(ctx context.Context, sessionID string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.session, sessionID)

	return nil
}

// GetSessionUserID implements Store.
func (m *MemoryStore) GetSessionUserID(ctx context.Context, sessionID string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	m.mu.RLock()
	defer m.mu.Unlock()

	return m.session[sessionID], nil
}

// GetUserByEmail implements Store.
func (m *MemoryStore) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	if err := ctx.Err(); err != nil {
		return domain.User{}, err
	}

	m.mu.RLock()
	defer m.mu.Unlock()

	return m.userByEmail[email], nil
}

// GetUserByID implements Store.
func (m *MemoryStore) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	if err := ctx.Err(); err != nil {
		return domain.User{}, err
	}

	m.mu.RLock()
	defer m.mu.Unlock()

	return m.usersByID[id], nil
}

// in Go, interfaces are already reference-like containers. A pointer to an interface is almost never needed.
func NewMemoryStore() Store {
	return &MemoryStore{
		usersByID:     make(map[string]domain.User),
		userByEmail:   make(map[string]domain.User),
		passwordHashs: make(map[string]string),
		session:       make(map[string]string),
	}
}
