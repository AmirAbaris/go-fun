package main

import "fmt"

// TODO(phase1-03): Learn methods and receivers.
//
// Task:
//   1. Define struct Book with Title and Author.
//   2. Implement String() string on *Book (fmt.Stringer interface).
//   3. Add a method with value receiver that does NOT mutate — explain why.
//   4. Add a method with pointer receiver that mutates a field — explain why pointer.
//
// Learn: https://go.dev/tour/methods/1 — methods, https://go.dev/tour/methods/9 — Stringer.
// Run: go run ./exercises/03_methods

func main() {
	var _ fmt.Stringer
	// your code here
}
