package main

import (
	"fmt"
	"sync"
	"time"
)

// worker simulates a task done by a goroutine.
// - id: the worker number
// - wg: WaitGroup to signal completion
// - ch: channel to send results
func worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // mark this worker as done when function exits

	// Simulate some work with sleep
	time.Sleep(time.Second)

	// Send result message into the channel
	ch <- fmt.Sprintf("Worker %d finished work", id)
}

func main() {
	var wg sync.WaitGroup      // WaitGroup to wait for all goroutines
	ch := make(chan string, 3) // Buffered channel with capacity 3

	// Launch 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)             // increment WaitGroup counter
		go worker(i, &wg, ch) // start worker goroutine
	}

	wg.Wait() // Wait until all workers have called wg.Done()
	close(ch) // Close the channel so range loop knows no more values will come

	// Read all results from channel
	for msg := range ch {
		fmt.Println(msg)
	}
}
