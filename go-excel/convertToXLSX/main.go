package main

import (
	"fmt"
	"os/exec"

	"github.com/xuri/excelize/v2"
)

func main() {
	arg1 := ""
	arg2 := ""

	err := convertToXlsx(arg1, arg2)
	if err != nil {
		fmt.Println("Error executing conversion:", err)
	}

	// Excelで操作
	f, err := excelize.OpenFile(arg2)
	if err != nil {
		fmt.Println("error opening excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	// 指定したシートの指定したセルの値を取得
	value, err := f.GetCellValue(sheetName, "A6")
	if err != nil {
		fmt.Println(fmt.Errorf("error getting cell value: %w", err))
	}

	fmt.Println(value)

	// 例として、(1, 1)をA1形式に変換
	cellRef := CoordsToCellRef(1, 1)
	fmt.Println(cellRef) // 出力: A1
}

// rowIndexとcolIndexからA1形式のセルアドレスに変換する関数
func CoordsToCellRef(colIndex, rowIndex int) string {
	col, _ := excelize.ColumnNumberToName(colIndex)
	return fmt.Sprintf("%s%d", col, rowIndex)
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
