# Phase 2b — Authentication API (stdlib + bcrypt)

Build user registration, login, and protected routes.

```bash
go run ./cmd/server
```

## Endpoints (you implement)

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/register` | No | Create account |
| POST | `/login` | No | Authenticate |
| POST | `/logout` | Yes | End session |
| GET | `/me` | Yes | Current user profile |

## Auth strategy

Pick **one** and implement fully before the other:
- **Session cookie** (recommended first) — store session ID server-side, set HttpOnly cookie
- **Bearer JWT** — sign token with HMAC, validate in middleware

The TODO in `internal/middleware/auth.go` guides this choice.

## Allowed non-stdlib dependency

- `golang.org/x/crypto/bcrypt` — password hashing only

## Learn

- Password hashing (never store plaintext)
- Auth middleware + context values for user ID
- `crypto/rand` for session tokens
