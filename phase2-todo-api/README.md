# Phase 2a — Todo API (stdlib only)

Build a REST Todo API using only the Go standard library.

```bash
go run ./cmd/server
```

## Endpoints (you implement)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Health check |
| GET | `/todos` | List all todos |
| POST | `/todos` | Create todo |
| GET | `/todos/{id}` | Get one todo |
| PATCH | `/todos/{id}` | Update todo |
| DELETE | `/todos/{id}` | Delete todo |

## Architecture

```
cmd/server          → wire dependencies, start server
internal/config     → env config
internal/domain     → Todo struct, validation (no net/http)
internal/store      → Store interface + in-memory impl
internal/handler    → HTTP handlers, JSON encode/decode
internal/middleware → logging, recover, request ID
```

## Learn

- `net/http`, `encoding/json`
- Middleware chaining
- Dependency injection via constructor params (no globals)

## Stdlib docs

- https://pkg.go.dev/net/http
- https://pkg.go.dev/encoding/json
