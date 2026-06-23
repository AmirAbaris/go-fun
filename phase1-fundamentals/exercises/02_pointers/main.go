package main

import "fmt"

// TODO(phase1-02): Learn pointers.
//
// Task:
//   1. Create a function that tries to increment an int by value — observe it doesn't work.
//   2. Rewrite it to take *int and increment successfully.
//   3. Demonstrate & (address-of) and * (dereference).
//   4. Show what happens when you dereference a nil pointer (and how to guard against it).
//
// Learn: https://go.dev/tour/moretypes/1 — pointers.
// Run: go run ./exercises/02_pointers

func Increase(number *int) {
	*number++
}

func main() {
	num := 0
	
	fmt.Println("Before: ", num)

	Increase(&num)

	fmt.Println("After: ", num)
}
