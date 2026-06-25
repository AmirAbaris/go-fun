package main

import (
	"context"
	"fmt"
	"time"
)

// TODO(phase1-08): Learn context.Context.
//
// Task:
//   1. Write fakeDBQuery(ctx context.Context, delay time.Duration) error that respects cancellation.
//   2. Call it with context.WithTimeout — verify it returns when timeout fires.
//   3. Call it with context.WithCancel — cancel manually and verify early return.
//   4. Pass ctx as the first parameter (Go convention).
//
// Learn: https://pkg.go.dev/context, https://go.dev/blog/context
// Run: go run ./exercises/08_context

func fakeDBQuery(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		fmt.Println("query finished successfully")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// your code here
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := fakeDBQuery(ctx, 2*time.Second); err != nil {
		fmt.Println("error: ", err)
	}
}
