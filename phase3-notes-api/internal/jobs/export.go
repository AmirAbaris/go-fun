package jobs

import (
	"context"

	"github.com/amirabaris/go-fun/phase3-notes-api/internal/queue"
	"github.com/amirabaris/go-fun/phase3-notes-api/internal/service"
)

// TODO(phase3-jobs): Implement background job handlers.
//
// Milestone 6 — worker calls these after dequeuing.
//
// Task:
//   1. ExportHandler holds *service.NoteService.
//   2. Handle(ctx, job queue.Job) error — dispatch on job.Type.
//   3. handleExportNotes:
//        decode payload for user_id
//        call service to list all notes for user
//        write JSON to ./exports/{user_id}.json (or log "would export N notes")
//   4. Return errors for retry logic (worker can log and continue).
//   5. Keep job handlers idempotent where possible.
//
// Learn: background jobs — API enqueues, worker processes asynchronously.

type ExportHandler struct {
	svc *service.NoteService
}

func NewExportHandler(svc *service.NoteService) *ExportHandler {
	panic("TODO: implement NewExportHandler")
}

func (h *ExportHandler) Handle(ctx context.Context, job queue.Job) error {
	panic("TODO: implement Handle — dispatch on job.Type")
}
