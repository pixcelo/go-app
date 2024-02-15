package main

import "fmt"

// Goは基本的に、値渡し（ローカルコピー） (メモリ内の新しい変数)
// func main() {
//     firstName := "John"
//     updateName(firstName)
//     fmt.Println(firstName)
// }

// func updateName(name string) {
//     name = "David"
// }

func main() {
	firstName := "John"

	// ポインター & をつけると参照渡しになる
	updateName(&firstName)
	fmt.Println(firstName)
}

// ポインターを受け取るには * （Refみたいな感じ）
// その後にあるオブジェクトのアドレスを受け取る
func updateName(name *string) {
	*name = "David"
}
