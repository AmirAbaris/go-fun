package queue

import "context"

// TODO(phase3-queue): Implement job queue abstraction.
//
// Milestone 6 — start with in-memory, optionally upgrade later.
//
// Task:
//   1. Define Job struct: Type (string), Payload ([]byte or typed).
//   2. Define Enqueuer interface: Enqueue(ctx, job Job) error.
//   3. Define Dequeuer interface: Dequeue(ctx) (Job, error) — blocks or polls.
//   4. MemoryQueue — channel-backed implementation for local dev.
//   5. Export job type constant: JobTypeExportNotes = "export_notes".
//   6. Export payload: { "user_id": "..." }.
//
// Optional milestone 7:
//   Replace MemoryQueue with Redis Streams or NATS — same interface, swap in cmd/.
//
// Learn: queue decouples API request from long-running work.

const JobTypeExportNotes = "export_notes"

type Job struct {
	Type    string
	Payload []byte
}

type Enqueuer interface {
	Enqueue(ctx context.Context, job Job) error
}

type Dequeuer interface {
	Dequeue(ctx context.Context) (Job, error)
}

func NewMemoryQueue(buffer int) (Enqueuer, Dequeuer) {
	panic("TODO: implement NewMemoryQueue with buffered channel")
}
