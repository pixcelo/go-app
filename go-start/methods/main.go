package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs メソッドは v という名前の Vertex 型のレシーバを持つ
// 構造体V　Vertex型 に Absメソッドを関連付ける例
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 5
}
