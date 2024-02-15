package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

// 引数に名前をつけて関数内で代入 最後にreturnとシンプルに書ける
func sum2(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}

// https://pkg.go.dev/strconv
