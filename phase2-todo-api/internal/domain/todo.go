package domain

import "time"

// TODO(phase2-todo-domain): Define the Todo domain model.
//
// Task:
//   1. Define Todo struct: ID (string), Title (string), Done (bool), CreatedAt (time.Time).
//   2. Define CreateTodoInput with Title field.
//   3. Define UpdateTodoInput with optional Title (*string) and Done (*bool).
//   4. Implement Validate() error on CreateTodoInput — title must be non-empty.
//   5. Keep this package free of net/http imports (pure domain).
//
// Learn: domain-driven design basics — validation lives with the type.

type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
}

type CreateTodoInput struct {
	Title string
}

type UpdateTodoInput struct {
	Title *string
	Done  *bool
}

func (in CreateTodoInput) Validate() error {
	panic("TODO: implement validation")
}
