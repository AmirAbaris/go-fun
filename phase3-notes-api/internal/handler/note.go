package handler

import (
	"net/http"

	"github.com/amirabaris/go-fun/phase3-notes-api/internal/service"
)

// TODO(phase3-handler): Implement HTTP handlers for notes.
//
// Milestone 4 — transport layer only.
//
// Task:
//   1. NoteHandler holds *service.NoteService (injected).
//   2. Handlers:
//        ListNotes  GET    /notes
//        CreateNote POST   /notes
//        GetNote    GET    /notes/{id}
//        UpdateNote PATCH  /notes/{id}
//        DeleteNote DELETE /notes/{id}
//        ExportNotes POST /notes/export
//   3. Read userID from context (auth middleware from phase2 pattern).
//   4. Decode/encode JSON DTOs — don't expose domain internals directly if shapes differ.
//   5. Map domain errors to HTTP status: NotFound→404, Forbidden→403, Invalid→400.
//   6. NO sql, pgx, or redis imports in this package.
//
// Learn: handler is thin — validate HTTP, call service, map response.

type NoteHandler struct {
	svc *service.NoteService
}

func NewNoteHandler(svc *service.NoteService) *NoteHandler {
	panic("TODO: implement NewNoteHandler")
}

func (h *NoteHandler) ListNotes(w http.ResponseWriter, r *http.Request) {
	panic("TODO: implement ListNotes")
}
