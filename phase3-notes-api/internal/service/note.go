package service

import (
	"context"

	"github.com/amirabaris/go-fun/phase3-notes-api/internal/domain"
	"github.com/amirabaris/go-fun/phase3-notes-api/internal/repository"
)

// TODO(phase3-service): Implement business logic layer.
//
// Milestone 3 — orchestrates repository, cache, and queue.
//
// Task:
//   1. NoteService struct with injected deps:
//        noteRepo repository.NoteRepository
//        cache    NoteCacher (interface — define in this file or cache package)
//        queue    Enqueuer (interface — define in queue package)
//   2. NewNoteService(...) *NoteService.
//   3. Methods:
//        CreateNote(ctx, userID, input) — validate, repo.Create, invalidate cache
//        GetNote(ctx, userID, noteID) — check cache first, then repo
//        ListNotes(ctx, userID, tagFilter string)
//        UpdateNote(ctx, userID, noteID, input) — invalidate cache on success
//        DeleteNote(ctx, userID, noteID) — invalidate cache
//        EnqueueExport(ctx, userID) — push ExportJob to queue
//   4. Return domain errors, not wrapped infrastructure errors (map pgx errors here).
//   5. NO net/http imports — service doesn't know about HTTP.
//
// Learn: clean architecture — application/use-case layer.

type NoteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) *NoteService {
	panic("TODO: implement NewNoteService with cache and queue deps")
}

func (s *NoteService) CreateNote(ctx context.Context, userID string, input domain.CreateNoteInput) (domain.Note, error) {
	panic("TODO: implement CreateNote")
}
