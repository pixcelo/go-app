package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run .\command.go 3 5
// Sum: 8

// go run .
// panic: runtime error: index out of range [1] with length 1

func main() {
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum:", number1+number2)
}
