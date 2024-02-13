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
	// num1 := 123
	// var num2 int = 12345678
	// num3 := 1.23
	// var num4 float64 = 1.2233
	// fmt.Println(num1)
	// fmt.Println(num2)
	// fmt.Println(num3)
	// fmt.Println(num4)
	// fmt.Println(reflect.TypeOf((num1))) // int
	// fmt.Println(reflect.TypeOf((num2))) // int
	// fmt.Println(reflect.TypeOf((num3))) // float64
	// fmt.Println(reflect.TypeOf((num4))) // float64

	// String
	// var string1 string = "hello"
	// string2 := "good morning"
	// fmt.Println(string1)
	// fmt.Println(reflect.TypeOf(string1))
	// fmt.Println(string2)
	// fmt.Println(reflect.TypeOf(string2))

	// Boolean
	a := 10
	b := 1
	var num_bool bool = a > b
	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))

	// Array
	arry := [3]string{"Tom", "Ken", "Mary"}
	arry[1] = "Jim"

	fmt.Println(arry[0])
	fmt.Println(arry[1])
	fmt.Println(reflect.TypeOf(arry))

	// 省略記法
	arryC := [...]string{"Tom", "Ken", "Mary"}
	fmt.Println(arryC[0])

	arryD := [2][2]string{{"1", "2"}, {"3", "4"}}
	fmt.Println(arryD[0][0])
	fmt.Println(arryD[0][1])
	fmt.Println(arryD[1][0])
	fmt.Println(arryD[1][1])
	fmt.Println(reflect.TypeOf(arryD))

	// age := 5
	// if age >= 2 {
	// 	fmt.Println("adult")
	// } else if age == 0 {
	// 	fmt.Println(("baby"))
	// } else {
	// 	fmt.Println(("child"))
	// }

	// 簡易文を使ったifステートメント(セミコロンで区切る)
	if age := 30; age >= 20 {
		fmt.Println("adult")
	} else if age == 0 {
		fmt.Println(("baby"))
	} else {
		fmt.Println(("child"))
	}

}
