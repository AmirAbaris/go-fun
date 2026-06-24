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

	// The range loop blocks waiting for values from the channel.
	//
	// Flow:
	// 1. Worker goroutines send their IDs into the channel.
	// 2. This loop receives those values and prints them.
	// 3. After a worker successfully sends its value, it finishes and calls wg.Done().
	// 4. When all workers are done, wg.Wait() in the closer goroutine unblocks.
	// 5. The closer goroutine closes the channel.
	// 6. The range loop detects that the channel is closed and exits automatically.
	//
	// Important: closing the channel does NOT unblock the workers.
	// Receiving from the channel unblocks the workers.
	// Closing the channel simply tells this loop that no more values will arrive.
	for id := range ch {
		fmt.Printf("id is: %d\n", id)
	}
}
