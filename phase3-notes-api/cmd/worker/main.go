package main

// TODO(phase3-worker-main): Background job consumer entry point.
//
// Milestone 6 — build after queue and jobs packages exist.
//
// Task:
//   1. Load config (same as API — DATABASE_URL, REDIS_URL, etc.).
//   2. Wire service layer (same deps as API minus HTTP handlers).
//   3. Create job handler: jobs.NewExportHandler(svc).
//   4. Poll queue.Dequeue in a loop (or subscribe to Redis stream / NATS).
//   5. On SIGINT/SIGTERM, drain in-flight jobs then exit.
//   6. Process ExportNotes job: fetch all user notes, write JSON to file or log stub.
//
// Run: go run ./cmd/worker

func main() {
	panic("TODO: implement worker composition root")
}
