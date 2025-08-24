package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func worker(id int, ch chan string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Duration(id) * 300 * time.Millisecond)
		ch <- fmt.Sprintf("Worker %d finished job %d", id, i)
	}
}

func main() {
	// Interfaces
	shapes := []Shape{
		Rectangle{
			Width:  10,
			Height: 5,
		},
		Circle{
			Radius: 7,
		},
	}

	for _, s := range shapes { // _ is index
		fmt.Println("Area: ", s.Area())
	}

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	person := map[string]string{"name": "Harshit", "role": "Developer"}
	file, _ := os.Create("person.json") // _ is error info
	defer file.Close()
	json.NewEncoder(file).Encode(person)
	fmt.Println("JSON written to person.json ✅")

	// Goroutines + Channels
	ch := make(chan string)
	go worker(1, ch)
	go worker(2, ch)

	// Use select to wait on multiple goroutines
	for i := 0; i < 6; i++ {
		select {
		case msg := <-ch:
			fmt.Println("Received:", msg)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout waiting for worker")
		}
	}

	fmt.Println("Program finished 🚀")
}
