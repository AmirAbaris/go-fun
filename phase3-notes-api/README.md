# Phase 3 — Notes API (Clean Architecture)

Authenticated users create notes with optional tags. Full stack: pgx, sqlc, migrations, Redis, background jobs.

## Milestones (build in order)

1. Migrations + sqlc
2. Repository layer
3. Service layer
4. Handler + middleware
5. Redis cache
6. Background worker
7. (Optional) Upgrade queue to Redis streams / NATS

## Run

```bash
# Apply migrations (after installing golang-migrate)
migrate -path migrations -database "$DATABASE_URL" up

# Generate sqlc code
sqlc generate

# API server
go run ./cmd/api

# Background worker
go run ./cmd/worker
```

## Layer rules

```
handler  → service  → repository (interface)
                ↓           ↓
              cache    postgres (sqlc)
                ↓
              queue → jobs/worker
```

- **domain** — no imports from handler, repository, cache, queue
- **service** — depends on repository interfaces, not pgx types
- **handler** — HTTP only, no SQL
- **repository/postgres** — sqlc-generated queries + domain mapping
- **cmd/** — composition root (wire everything here)

## Dependencies (add when implementing)

- `github.com/jackc/pgx/v5`
- `github.com/golang-migrate/migrate/v4`
- `github.com/redis/go-redis/v9`
- sqlc CLI: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

## Endpoints (you implement)

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/register` | No | Create account |
| POST | `/login` | No | Authenticate |
| GET | `/notes` | Yes | List user's notes |
| POST | `/notes` | Yes | Create note |
| GET | `/notes/{id}` | Yes | Get note (cache this) |
| PATCH | `/notes/{id}` | Yes | Update note |
| DELETE | `/notes/{id}` | Yes | Delete note |
| POST | `/notes/export` | Yes | Enqueue export job |
