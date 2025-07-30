package main

import "fmt"

// Basic Go syntax examples
func main() {
	// Variables
	var name string = "Go Learner"
	age := 25
	
	// Constants
	const pi = 3.14159
	
	// Basic types
	var (
		isLearning bool    = true
		score      float64 = 95.5
		count      int     = 100
	)
	
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Age: %d, Score: %.2f, Count: %d\n", age, score, count)
	fmt.Printf("Learning Go: %t, Pi: %.5f\n", isLearning, pi)
	
	// Arrays and slices
	var numbers [3]int = [3]int{1, 2, 3}
	slice := []string{"Go", "is", "awesome"}
	
	fmt.Printf("Array: %v\n", numbers)
	fmt.Printf("Slice: %v\n", slice)
	
	// Maps
	languages := map[string]int{
		"Go":     2009,
		"Python": 1991,
		"Java":   1995,
	}
	
	fmt.Printf("Languages: %v\n", languages)
	fmt.Printf("Go was created in: %d\n", languages["Go"])
}
