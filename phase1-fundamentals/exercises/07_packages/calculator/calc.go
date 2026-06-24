// Package calculator provides basic arithmetic operations.
package calculator

import "errors"

// TODO(phase1-07): Learn packages and exported identifiers.
//
// Task:
//   1. Implement exported functions Add, Subtract, Multiply, Divide.
//   2. Divide should return (float64, error) on division by zero.
//   3. Keep an unexported helper (e.g. validate) — confirm it can't be used from outside.
//   4. Add a package doc comment above `package calculator`.
//
// Learn: https://go.dev/tour/basics/3 — exported names.
// Used by: exercises/07_packages/cmd/demo/main.go

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if err := validate(b); err != nil {
		return 0, err
	}
	return a / b, nil
}

func validate(b float64) error {
	if b == 0 {
		return errors.New("a and b must be non-zero")
	}
	return nil
}
