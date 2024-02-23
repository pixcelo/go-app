package main

import "fmt"

// マップ　Go のマップは基本的にハッシュ テーブルであり、キーと値のペアのコレクション
// 初期化 map[T]T
func fn1() {
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)
	fmt.Println(studentsAge["john"])
	fmt.Println(studentsAge["christy"]) // keyが存在しない場合、panicにならず　規定値が入る intなら0

	// 初期化しない場合は　studentsAge := make(map[string]int)
	// C#でいうと var studentsAge = new Dictionary<string, int>();
	// studentsAge := make(map[string]int)
	// studentsAge["john"] = 32
	// studentsAge["bob"] = 31
}

// マップの存在確認は、第２返り値でbool
func fn2() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
}

// 項目の削除 delete(map, key)
func fn3() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	delete(studentsAge, "john")
	fmt.Println(studentsAge)
}

// マップのループ　for eachの代わり
func fn4() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func main() {
	fn4()
}
