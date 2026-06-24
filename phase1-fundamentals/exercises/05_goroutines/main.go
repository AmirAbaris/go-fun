package main

import (
	"fmt"
	"sync"
)

// TODO(phase1-05): Learn goroutines and sync.WaitGroup.
//
// Task:
//   1. Launch 5 goroutines that each print their worker id.
//   2. Use sync.WaitGroup so main waits for all workers.
//   3. Launch goroutines that increment a shared counter WITHOUT a mutex — run with:
//      go run -race ./exercises/05_goroutines
//      Observe the race detector warning (you will fix this in exercise 06).
//
// Learn: https://go.dev/tour/concurrency/1, https://pkg.go.dev/sync#WaitGroup
// Run: go run ./exercises/05_goroutines

func Runner(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker id is: %d\n", id)
}

func main() {
	// your code here
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go Runner(i, &wg)
	}

	wg.Wait()
}
