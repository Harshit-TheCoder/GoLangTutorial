package main

import (
	"context"
	"fmt"
	"time"
)

// worker simulates a background task that keeps running until its context is cancelled.
func worker(ctx context.Context) {
	for {
		select {
		// <-ctx.Done() is a signal channel that closes when the context is cancelled or times out
		case <-ctx.Done():
			// Once context is cancelled, worker prints why it stopped and exits the loop
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			// If context is still active, do some "work"
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) // simulate time taken for work
		}
	}
}

func main() {
	// Create a context that will automatically cancel after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // ensure cancel is called to free resources

	// Start worker in a separate goroutine, passing the context
	go worker(ctx)

	// Main function waits for 3 seconds
	// Worker will run for ~2 seconds and then stop when the context times out
	time.Sleep(3 * time.Second)

	fmt.Println("Main finished")
}
