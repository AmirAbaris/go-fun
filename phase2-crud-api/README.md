# Phase 2c — CRUD API + PostgreSQL

Build a Product CRUD API using `database/sql` and PostgreSQL.

```bash
# Start Postgres locally, then:
export DATABASE_URL="postgres://user:pass@localhost:5432/products?sslmode=disable"
go run ./cmd/server
```

## Endpoints (you implement)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Health check (include DB ping) |
| GET | `/products` | List products |
| POST | `/products` | Create product |
| GET | `/products/{id}` | Get one product |
| PUT | `/products/{id}` | Update product |
| DELETE | `/products/{id}` | Delete product |

## Setup

1. Run `migrations/001_init.sql` against your database manually (or use psql).
2. Phase 3 will introduce `golang-migrate` for versioned migrations.

## Learn

- `database/sql` — Open, Ping, QueryContext, ExecContext
- Parameterized queries (`$1`, `$2`) — never string-concatenate SQL
- `context.Context` on every DB call
- Connection lifecycle in main

## Allowed dependency

- `github.com/lib/pq` — Postgres driver for database/sql
