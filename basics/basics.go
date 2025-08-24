package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int {
	return a + b
}

type Person struct {
	Name string
	Age  int
}

func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func squares(nums []int, ch chan int) {
	for _, n := range nums {
		ch <- n * n
	}
	close(ch)
}

func main() {
	// --- Variables ---
	var x int = 10
	y := 20
	fmt.Println("x + y =", x+y)

	// --- Arrays ---
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// --- Slices ---
	slice := []string{"Go", "is", "fun"}
	slice = append(slice, "!")
	fmt.Println("Slice:", slice)

	// --- Maps ---
	m := make(map[string]int)
	m["apples"] = 5
	m["oranges"] = 3
	fmt.Println("Map:", m)

	// --- Loops ---
	for i := 1; i <= 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- Conditionals ---
	if x > y {
		fmt.Println("x is greater")
	} else {
		fmt.Println("y is greater")
	}

	// --- Functions ---
	sum := add(5, 7)
	fmt.Println("Sum:", sum)

	// --- Structs & Methods ---
	p := Person{
		Name: "Harshit",
		Age:  20,
	}
	p.greet()

	// --- Pointers ---
	num := 42
	ptr := &num
	fmt.Println("Pointer value:", *ptr)

	// --- Goroutines ---
	fmt.Println("\n--- Goroutines Example ---")
	go printNumbers() //runs concurrently with main

	// While goroutine runs, main can do other work
	for i := 1; i <= 3; i++ {
		fmt.Println("Main function working...")
		time.Sleep(700 * time.Millisecond)
	}

	// --- Channels --- send/receive values safely between goroutines
	fmt.Println("\n--- Channels Example ---")
	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	go squares(numbers, ch)

	for sq := range ch { // recieve until channel closed
		fmt.Println("Square: ", sq)
	}

	fmt.Println("\n Program finished ✅")

}
