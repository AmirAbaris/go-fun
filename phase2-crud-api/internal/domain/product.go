package domain

import "time"

// TODO(phase2-crud-domain): Define Product domain model.
//
// Task:
//   1. Product struct: ID (int64), Name, Description, Price (int cents or float64 — pick one), CreatedAt.
//   2. CreateProductInput, UpdateProductInput with Validate().
//   3. Name required, Price must be > 0.
//   4. No database/sql or net/http imports in this package.

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int64 // cents — avoids float money bugs
	CreatedAt   time.Time
}

type CreateProductInput struct {
	Name        string
	Description string
	Price       int64
}

func (in CreateProductInput) Validate() error {
	panic("TODO: implement validation")
}

type UpdateProductInput struct {
	Name        *string
	Description *string
	Price       *int64
}
