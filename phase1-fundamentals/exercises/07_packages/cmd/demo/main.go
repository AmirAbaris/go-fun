package main

import (
	"fmt"

	"github.com/amirabaris/go-fun/phase1-fundamentals/exercises/07_packages/calculator"
)

// TODO(phase1-07-demo): Import and use the calculator package.
//
// Task:
//   1. Call each exported calculator function and print results.
//   2. Handle the error from Divide.
//   3. Notice the import path matches your go.mod module path.
//
// Run: go run ./exercises/07_packages/cmd/demo

func main() {
	fmt.Println(calculator.Add(1, 2))
	fmt.Println(calculator.Divide(1, 2))
	fmt.Println(calculator.Multiply(1, 2))
	fmt.Println(calculator.Subtract(1, 2))
}
