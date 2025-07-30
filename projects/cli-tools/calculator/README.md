# Calculator CLI Tool

A simple command-line calculator built in Go.

## Usage

```bash
go run main.go <num1> <operator> <num2>
```

## Examples

```bash
go run main.go 5 + 3
# Output: 5.00 + 3.00 = 8.00

go run main.go 10 - 4
# Output: 10.00 - 4.00 = 6.00

go run main.go 7 * 8
# Output: 7.00 * 8.00 = 56.00

go run main.go 15 / 3
# Output: 15.00 / 3.00 = 5.00
```

## Supported Operations

- Addition (+)
- Subtraction (-)
- Multiplication (*)
- Division (/)

## Error Handling

- Invalid numbers
- Division by zero
- Invalid operators
- Incorrect number of arguments

## Learning Goals

This project demonstrates:
- Command-line argument parsing
- String to number conversion
- Error handling
- Switch statements
- Basic arithmetic operations
