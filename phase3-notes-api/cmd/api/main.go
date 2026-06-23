package main

// TODO(phase3-api-main): Composition root — wire all dependencies and start API server.
//
// Milestone 4+ — build this after service and handler layers exist.
//
// Task:
//   1. Load config from internal/config.
//   2. Connect pgx pool: pgxpool.New(ctx, databaseURL).
//   3. Run migrations (or assume already applied).
//   4. Wire layers:
//        repo := postgres.NewNoteRepository(pool)
//        cache := redis.NewNoteCache(redisClient)  // milestone 5
//        svc  := service.NewNoteService(repo, cache, queue)
//        h    := handler.NewNoteHandler(svc)
//   5. Register routes with auth middleware (reuse patterns from phase2-auth-api).
//   6. Graceful shutdown: pool.Close(), redis.Close(), server.Shutdown.
//
// Run: go run ./cmd/api

func main() {
	panic("TODO: implement API composition root")
}
