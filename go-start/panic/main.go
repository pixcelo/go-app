/*
 範囲外のインデックスを使用して配列にアクセスしようとしたり、
 il ポインターを逆参照したりして、実行時エラーが発生すると、Go プログラムはパニック状態になる

 https://learn.microsoft.com/ja-jp/training/modules/go-control-flow/4-use-defer-statement
*/

package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

// func main() {
// 	highlow(2, 0)
// 	fmt.Println(("Program finished successfully!"))
// }

// recover 例外処理
func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
