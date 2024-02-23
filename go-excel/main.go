package main

import (
	"fmt"
	"os/exec"
)

func main() {
	arg1 := ""
	arg2 := ""

	err := convertToXlsx(arg1, arg2)
	if err != nil {
		fmt.Println("Error executing conversion:", err)
	}
}

// Convert xls to xlsx
func convertToXlsx(sourcePath, targetPath string) error {
	path := "./net6.0/ExcelConverter.exe"
	cmd := exec.Command(path, sourcePath, targetPath)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %w", err)
	}

	return nil
}
