package main

import (
	"fmt"
	"os/exec"
)

func main() {
	exePath := "./net6.0/ExcelConverter.exe"

	arg1 := ""
	arg2 := ""

	cmd := exec.Command(exePath, arg1, arg2)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// 出力を表示
	fmt.Println("Command output:", string(output))
}
