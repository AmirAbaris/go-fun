package main

import (
	"fmt"
	"sync"
)

// TODO(phase1-06): Learn channels and select.
//
// Task:
//   1. Create an unbuffered channel — send/receive between two goroutines.
//   2. Create a buffered channel — show non-blocking send until full.
//   3. Fan-in: merge two channels into one using select.
//   4. Use context.WithCancel to stop workers gracefully.
//   5. Fix the race from exercise 05 using a channel or sync.Mutex.
//
// Learn: https://go.dev/tour/concurrency/2, https://gobyexample.com/select
// Run: go run ./exercises/06_channels

func runner(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- id
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go runner(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		// close the channel only when you are sure that no more values will be sent.
		close(ch)
	}()

	for id := range ch {
		fmt.Printf("id is: %d\n", id)
	}
}
