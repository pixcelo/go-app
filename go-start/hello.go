package main

//import "fmt"
import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println("Hello, world")

	// var num int
	// num = 1
	// num := 1
	// fmt.Println(num)

	// Number
	num1 := 123
	var num2 int = 12345678
	num3 := 1.23
	var num4 float64 = 1.2233
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(reflect.TypeOf((num1))) // int
	fmt.Println(reflect.TypeOf((num2))) // int
	fmt.Println(reflect.TypeOf((num3))) // float64
	fmt.Println(reflect.TypeOf((num4))) // float64
}
