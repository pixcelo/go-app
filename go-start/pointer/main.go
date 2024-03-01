package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func fn1() {
	// ポインタ型の変数を宣言
	// pがポインタ変数
	// *Personがポインタ型
	var p *Person

	p = &Person{
		Name: "Tom",
		Age:  20,
	}

	fmt.Printf("p :%+v\n", p)
	fmt.Printf("変数pに格納されているアドレス: %p", p)
	// 変数pに格納されているアドレス: 0xc000010030%
}

func fn2() {
	// 値として、pに代入
	p := Person{
		Name: "太郎",
		Age:  20,
	}

	fmt.Printf("最初のp :%+v\n", p)

	p2 := p
	p2.Name = "二郎"
	p2.Age = 21
	// pではなく
	fmt.Printf("p2で二郎に書き換えを行なったはずのp :%+v\n", p) // 値コピーのためオリジナルのpは変更されない

	// &pで*Person(Personのポインタ型)を生成する
	// p3はpのアドレスが格納されている状態になる
	p3 := &p
	p3.Name = "二郎"
	p3.Age = 21

	fmt.Printf("p3で二郎に書き換えを行なったp :%+v\n", p)
}

func fn3() {
	name := "太郎"
	fmt.Printf("name :%v\n", name)

	namePoint := &name

	// namePointは、&nameが格納されているだけなので、stringへのポインタである *string型の値が格納されている。
	fmt.Printf("namePoint :%v\n", namePoint)

	// namePointが指している変数は、"*namePoint"という感じで、"*"をつけて表す。
	fmt.Printf("namePoint :%v\n", *namePoint)
}

func main() {
	fn3()
}
