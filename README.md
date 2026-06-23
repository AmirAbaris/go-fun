# Go Practice Project

A guided, hands-on path to learning Go — from language fundamentals to production-style APIs.

**Rule:** Read the [standard library docs](https://pkg.go.dev/std) before reaching for third-party libraries in Phase 2.

## Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

## Phases

| Phase | Directory | Focus | Time |
|-------|-----------|-------|------|
| 1 | `phase1-fundamentals/` | Variables, structs, pointers, methods, errors, goroutines, channels, packages, context, testing | 1–2 weeks |
| 2a | `phase2-todo-api/` | `net/http`, JSON, middleware, DI (stdlib only) | ~1 week |
| 2b | `phase2-auth-api/` | Password hashing, sessions/JWT, auth middleware | ~1 week |
| 2c | `phase2-crud-api/` | `database/sql`, PostgreSQL, parameterized queries | ~1 week |
| 3 | `phase3-notes-api/` | Clean architecture, pgx, sqlc, migrations, Redis, background jobs | 2+ weeks |

## How to run each project

```bash
# Phase 1 — run an exercise
cd phase1-fundamentals
go run ./exercises/01_variables_structs
go test ./...

# Phase 2 — start an API server
cd phase2-todo-api && go run ./cmd/server
cd phase2-auth-api && go run ./cmd/server
cd phase2-crud-api  && go run ./cmd/server

# Phase 3 — API + worker
cd phase3-notes-api && go run ./cmd/api
cd phase3-notes-api && go run ./cmd/worker
```

Each subdirectory has its own `go.mod`. Initialize modules on first use:

```bash
cd phase1-fundamentals && go mod tidy
```

## Suggested weekly flow

| Week | Focus |
|------|-------|
| 1 | Phase 1 exercises 01–05 + stdlib reading |
| 2 | Phase 1 exercises 06–09; start Phase 2 Todo API |
| 3 | Finish Todo + Auth APIs |
| 4 | CRUD API + PostgreSQL |
| 5–6 | Phase 3 Notes API layer by layer |

## Exit criteria

### Phase 1
- [ ] `go test ./...` passes in `phase1-fundamentals`
- [ ] Can explain pointer vs value receivers
- [ ] Comfortable with `if err != nil` and error wrapping
- [ ] Know when to use goroutines vs channels

### Phase 2
- [ ] All three APIs run locally
- [ ] Middleware chain works (logging, recover, auth where applicable)
- [ ] Consistent JSON error responses
- [ ] CRUD API uses parameterized SQL (`$1`, `$2`, …)

### Phase 3
- [ ] Handler layer does not import `pgx`
- [ ] Migrations are versioned and reversible
- [ ] Redis cache invalidates on write
- [ ] Worker processes at least one job type

## How to use the TODOs

Every stub file starts with a `// TODO:` block describing:
- What to build
- Which Go concepts to learn
- How to run and verify your work

Implement the task, delete or check off the TODO when done, then move to the next file.
