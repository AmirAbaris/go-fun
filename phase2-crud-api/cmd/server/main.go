package main

// TODO(phase2-crud-main): Wire database and start HTTP server.
//
// Task:
//   1. Load config (PORT, DATABASE_URL).
//   2. Open *sql.DB with sql.Open("postgres", dsn) — import _ "github.com/lib/pq".
//   3. PingContext with timeout to verify connection.
//   4. defer db.Close().
//   5. Create store: store.NewPostgresStore(db).
//   6. Wire handler and middleware (same pattern as Todo API).
//   7. GET /health should call db.PingContext and return 503 if DB is down.
//   8. Graceful shutdown.
//
// Run: DATABASE_URL=... go run ./cmd/server

func main() {
	panic("TODO: implement server wiring with *sql.DB injection")
}
