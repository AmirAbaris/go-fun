package main

import "testing"

// TODO(phase1-01-test): Write a table-driven test for a small pure function you create.
//
// Task:
//   1. Extract a pure function from main (e.g. IsAdult(age int) bool).
//   2. Write table-driven tests with subtests (t.Run).
//   3. Cover at least: zero value edge case, typical case, boundary case.
//
// Learn: https://go.dev/doc/tutorial/add-a-test
// Verify: go test ./exercises/01_variables_structs -v

func TestGreet(t *testing.T) {
	person := Person{Name: "Abaris", Age: 23}
	got := person.Greet()
	want := "Hi Abaris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	
}
