package main

import "fmt"

// with an unbuffered channel, send and receive must happen from different goroutines.
func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	newValue := <-ch

	fmt.Print(newValue)
}
