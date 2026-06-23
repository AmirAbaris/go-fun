package cache

import (
	"context"

	"github.com/amirabaris/go-fun/phase3-notes-api/internal/domain"
)

// TODO(phase3-cache): Implement Redis caching for notes.
//
// Milestone 5 — cache-aside pattern.
//
// Task:
//   1. Define NoteCacher interface:
//        Get(ctx, noteID string) (domain.Note, bool, error)  // found bool
//        Set(ctx, note domain.Note, ttl time.Duration) error
//        Delete(ctx, noteID string) error
//   2. RedisNoteCache implements NoteCacher using go-redis.
//   3. Serialize note as JSON for Redis values.
//   4. Key pattern: "note:{id}".
//   5. Service calls cache on GetNote; invalidates on Update/Delete.
//   6. Handle cache miss gracefully — fall through to repository.
//
// Learn: https://redis.io/docs/manual/patterns/cache-aside/
// Dep: github.com/redis/go-redis/v9

type NoteCacher interface {
	Get(ctx context.Context, noteID string) (domain.Note, bool, error)
	Set(ctx context.Context, note domain.Note, ttl interface{}) error
	Delete(ctx context.Context, noteID string) error
}

func NewRedisNoteCache(redisURL string) (NoteCacher, error) {
	panic("TODO: implement NewRedisNoteCache")
}
