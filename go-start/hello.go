package main

import "fmt"

//import "fmt"
//import "reflect"

// "reflect"

// func sayHello() {
// 	fmt.Println("hello")
// }
func sayHello(greeting string) {
	fmt.Println(greeting)
}

// 同じデータ型は省略できる
// func cal(x int, y int) {
func cal(x, y int) {
	fmt.Println(x * y)
}

//戻り値があれば戻り値の型も必要
func cal2(x, y int) int {
	return x * y
}

// Pythonのように戻り値を複数返せる
func cal3(x, y int) (int, int) {
	return (x / y), (x * y)
}

func cal4(x, y int) (int, int) {
	a := x / y
	b := x * y
	return a, b
}

// 戻り値を宣言している場合、returnの引数は省略可能
func cal5(x, y int) (a int, b int) {
	a = x / y
	b = x * y
	return
}

// 構造体
type Person struct {
	name string
	age  int
}

func main() {
	// 構造体の初期化
	var person Person
	person.name = "Tom"
	person.age = 20
	fmt.Println((person))

	// 構造体の初期化2
	person2 := Person{"Ken", 15}
	fmt.Println((person2))

	// 構造体の初期化3　フィールド名の指定
	person3 := Person{name: "Foo", age: 30}
	fmt.Println((person3))

	// cal(2, 3)
	// result := cal2(2, 3)
	// fmt.Println(result)
	// result1, result2 := cal3(3, 3)
	// fmt.Println(result1, result2)

	// 関数を変数に代入
	// hello := func(greeting string) {
	// 	fmt.Println(greeting)
	// }

	// hello("hello")

	// 無名関数 jsの即時関数みたいな書き方
	// func(greeting string) {
	// 	fmt.Println(greeting)
	// }("hello")

	// sayHello("hello")
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
	// a := 10
	// b := 1
	// var num_bool bool = a > b
	// fmt.Println(num_bool)
	// fmt.Println(reflect.TypeOf(num_bool))

	// Array
	// arry := [3]string{"Tom", "Ken", "Mary"}
	// arry[1] = "Jim"

	// fmt.Println(arry[0])
	// fmt.Println(arry[1])
	// fmt.Println(reflect.TypeOf(arry))

	// // 省略記法
	// arryC := [...]string{"Tom", "Ken", "Mary"}
	// fmt.Println(arryC[0])

	// arryD := [2][2]string{{"1", "2"}, {"3", "4"}}
	// fmt.Println(arryD[0][0])
	// fmt.Println(arryD[0][1])
	// fmt.Println(arryD[1][0])
	// fmt.Println(arryD[1][1])
	// fmt.Println(reflect.TypeOf(arryD))

	// age := 5
	// if age >= 2 {
	// 	fmt.Println("adult")
	// } else if age == 0 {
	// 	fmt.Println(("baby"))
	// } else {
	// 	fmt.Println(("child"))
	// }

	// 簡易文を使ったifステートメント(セミコロンで区切る)
	// if age := 30; age >= 20 {
	// 	fmt.Println("adult")
	// } else if age == 0 {
	// 	fmt.Println(("baby"))
	// } else {
	// 	fmt.Println(("child"))
	// }

	// for文1
	// for i := 0; i <= 4; i++ {
	// 	fmt.Println(i)
	// }

	// for文2
	// i := 0
	// for i <= 4 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for文3
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	if i == 4 {
	// 		break
	// 	}
	// 	i++
	// }
}
