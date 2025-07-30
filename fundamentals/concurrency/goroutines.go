package main

import (
	"fmt"
	"time"
)

// Goroutines and channels example
func main() {
	// Channel for communication
	messages := make(chan string)
	
	// Start a goroutine
	go func() {
		time.Sleep(1 * time.Second)
		messages <- "Hello from goroutine!"
	}()
	
	// Multiple goroutines with buffered channel
	numbers := make(chan int, 3)
	
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i * i
			fmt.Printf("Sent: %d\n", i*i)
		}
		close(numbers)
	}()
	
	// Receive from channel
	msg := <-messages
	fmt.Println("Received:", msg)
	
	// Receive from buffered channel
	for num := range numbers {
		fmt.Printf("Received number: %d\n", num)
	}
	
	// Channel with select
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "from c1"
	}()
	
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "from c2"
	}()
	
	// Select statement
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
