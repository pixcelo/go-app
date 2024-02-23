package main

import "fmt"

func fn1() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])        // 0
	fmt.Println(a[1])        // 10
	fmt.Println(a[len(a)-1]) // 0
}

func fn2() {
	//cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	//fmt.Println(cities)
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
}

// 省略記号を使用して、最後の位置の値のみを指定する方法
func fn3() {
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}

// 多次元配列
func fn4() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}

// https://learn.microsoft.com/ja-jp/training/modules/go-data-types/2-slices
// スライス
func fn5() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
}

// スライス演算子
func fn6() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
}

// appendしたときのスライスの容量は2倍されていくのでメモリの消費に注意
func fn7() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

// スライスの項目の削除
func fn8() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		// appendで削除する要素の前までのスライスと、削除する要素の後のスライスを結合
		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}
}

// スライスのコピーを作成 copyの場合、スライス自体の構造（ポインタ、長さ、容量）は新たに作成される
// 単に slice2 := letters[1:4] とすると参照渡しになる（lettersを変更するとslice2も変更される）
func fn9() {
	letters := []string{"A", "B", "C", "D", "E"}
	slice2 := make([]string, 3)
	fmt.Println("Before", slice2)
	copy(slice2, letters[1:4])
	fmt.Println("After", letters)
}

func main() {
	fn9()
}
