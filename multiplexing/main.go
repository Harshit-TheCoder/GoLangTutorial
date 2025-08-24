package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two unbuffered channels to send/receive strings
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: waits 1 second, then sends a message to ch1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	// Goroutine 2: waits 1 second, then sends a message to ch2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from Channel 2"
	}()

	// Loop to handle exactly 2 incoming messages (from ch1 and ch2)
	for i := 0; i < 2; i++ {
		select {
		// If ch1 has a value ready, receive it and print
		case msg1 := <-ch1:
			fmt.Println(msg1)

		// If ch2 has a value ready, receive it and print
		case msg2 := <-ch2:
			fmt.Println(msg2)

		// If neither channel is ready within 1.5s, timeout triggers
		case <-time.After(1500 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}
