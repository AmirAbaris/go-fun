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

type Book struct {
	Title string
	Author string
}

func (b *Book) String() string {
	return b.Title + " by " + b.Author
}

func (b *Book) AddSpace() {
	b.Title = b.Title + " "
}

// because we didnt used pointer reciver, it just reads and cant muatate here
func (b Book) Log() {
	fmt.Println("logged ", b.Title)
}

func main() {
	var _ fmt.Stringer = (*Book)(nil)
	// your code here
	book := Book{
		Title:  "1984",
		Author: "George Orwell",
	}

	book.AddSpace()
	book.Log()

	fmt.Println(book)
}
