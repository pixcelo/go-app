package main

import "fmt"

// defer ステートメント
func main() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("defered", -i)
		fmt.Println("regular", i)
	}
}

/* 実行結果

main() 関数で regular 値の出力が終了すると、すべての遅延呼び出しが実行される
遅れた呼び出しからの出力は、待ち行列から飛び出したように、順序が逆になる (後入れ、先だし)

regular 1
regular 2
regular 3
regular 4
defered -4
defered -3
defered -2
defered -1
*/

// package main

// import (
//     "io"
//     "os"
//     "fmt"
// )

// func main() {
//     newfile, error := os.Create("learnGo.txt")
//     if error != nil {
//         fmt.Println("Error: Could not create file.")
//         return
//     }
//     defer newfile.Close()

//     if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
// 	    fmt.Println("Error: Could not write to file.")
//         return
//     }

//     newfile.Sync()
// }
