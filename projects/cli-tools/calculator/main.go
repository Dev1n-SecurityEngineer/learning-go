package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <num1> <operator> <num2>")
		fmt.Println("Example: calculator 5 + 3")
		os.Exit(1)
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number\n", os.Args[1])
		os.Exit(1)
	}

	operator := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number\n", os.Args[3])
		os.Exit(1)
	}

	var result float64
	var validOp bool = true

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		validOp = false
	}

	if !validOp {
		fmt.Printf("Error: '%s' is not a valid operator. Use +, -, *, or /\n", operator)
		os.Exit(1)
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
