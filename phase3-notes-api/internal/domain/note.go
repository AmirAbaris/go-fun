package domain

import "time"

// TODO(phase3-domain-note): Define Note entity and domain errors.
//
// Task:
//   1. Note struct: ID, UserID, Title, Body, Tags []string, CreatedAt, UpdatedAt.
//   2. CreateNoteInput, UpdateNoteInput with Validate().
//   3. Domain errors (sentinel): ErrNotFound, ErrForbidden, ErrInvalidInput.
//   4. NO imports from: handler, repository, cache, queue, pgx, redis.
//
// This is the innermost layer — pure Go types and business rules.

type Note struct {
	ID        string
	UserID    string
	Title     string
	Body      string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateNoteInput struct {
	Title string
	Body  string
	Tags  []string
}

func (in CreateNoteInput) Validate() error {
	panic("TODO: implement validation")
}

type UpdateNoteInput struct {
	Title *string
	Body  *string
	Tags  []string
}
