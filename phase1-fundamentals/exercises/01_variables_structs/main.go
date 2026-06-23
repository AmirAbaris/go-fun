package main

// TODO(phase1-01): Learn variables, structs, and interfaces.
//
// Task:
//   1. Define a struct Person with Name (string) and Age (int).
//   2. Add a method Greet() string using a pointer receiver.
//   3. Define interface Greeter with Greet(); assign *Person to Greeter.
//   4. Print zero values vs initialized values for int, string, bool, slice.
//
// Learn: https://go.dev/tour/basics/1 — structs, methods, interfaces.
// Run: go run ./exercises/01_variables_structs
// Verify: go test ./...

import "fmt"

type Person struct {
	Name string
	Age int
}

// in Go Uppercased method name is for -> public
// lowercased method name is for -> private
func (p *Person) Greet() string {
	// in Go if we dont use pointers it donest let us to mutate original struct
	// it creates a copy of the struct and mutates that value, not our orignal one.
	// so thats why we can use pointers to directly mutate the struct values
	return "Hi " + p.Name
}

func main() {
	person := Person{Name: "Abaris",Age: 23}
	fmt.Println(person.Greet())

}
