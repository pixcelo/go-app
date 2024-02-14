package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//import "fmt"

// "reflect"

//	func sayHello() {
//		fmt.Println("hello")
//	}
func sayHello(greeting string) {
	fmt.Println(greeting)
}

// 同じデータ型は省略できる
// func cal(x int, y int) {
func cal(x, y int) {
	fmt.Println(x * y)
}

// 戻り値があれば戻り値の型も必要
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

type Student struct {
	name string
	age  int
}

// 構造体にメソッドを定義
func (s Student) printName() {
	fmt.Println(s.name)
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

	student := Student{name: "Make", age: 10}
	student.printName()

	fmt.Println(("====================="))

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

	// まとめて変数を宣言
	// var (
	// 	firstName string = "John"
	// 	lastName  string = "Doe"
	// 	age       int    = 32
	// )

	// 型推定
	// var (
	// 	firstName = "John"
	// 	lastName  = "Doe"
	// 	age       = 32
	// )

	// カンマ区切りの書き方
	// var (
	// 	firstName, lastName, age = "John", "Doe", 32
	// )

	// fmt.Println((firstName))
	// fmt.Println((lastName))
	// fmt.Println((age))

	// := は新規の変数宣言のみ使える(定義された場所がわかる)
	// firstName, lastName := "John", "Doe"
	// age := 32
	// fmt.Println(firstName, lastName, age)

	// 定数
	// const MAX_NUMBER = 200
	// fmt.Println(MAX_NUMBER)

	// rune は単に、int32 データ型の別名
	// rune := 'G'
	// fmt.Println(rune) // 71

	// var integer32 int = 2147483648
	// fmt.Println(integer32)

	// 文字列のエスケープ
	// \n 改行
	// \r 復帰
	// \t タブ
	// \' 単一引用符
	// \" 二重引用符
	// \\ バックスラッシュ
	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)

	// 変数を初期化しない場合、規定値が入る
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// 明示的に指定する型変換(暗黙的なキャストは出来ない)
	i, _ := strconv.Atoi("-42") // 文字列型を数値型に変換
	s := strconv.Itoa(-42)      // 数値型を文字列型に変換
	fmt.Println(i, s)
	fmt.Println(reflect.TypeOf((i)))
	fmt.Println(reflect.TypeOf((s)))

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
