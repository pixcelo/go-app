package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <table_name> <excel_file>")
	}

	// 引数から取得
	tableName := os.Args[1]
	excelPath := os.Args[2]
	// tableName := ""
	// excelPath := ""

	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		log.Fatalf("Error opening Excel file: %v", err)
	}

	// 1シート目のデータをすべて取得 [][]string
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		log.Fatalf("Error getting rows from sheet: %v", err)
	}

	// 1行目の物理名を取得　※今回は列数を使用
	columns := make(map[int]string)
	for idx, col := range rows[0] {
		columns[idx] = col
	}

	var values []string
	maxCols := len(columns)

	// 3行目から処理
	for _, row := range rows[2:] {
		var rowValues []string
		for i := 0; i < maxCols; i++ {
			// 値がなければNULLとして挿入、セル値は「'」で囲む
			if i < len(row) {
				rowValues = append(rowValues, "'"+strings.ReplaceAll(row[i], "'", "''")+"'")
			} else {
				rowValues = append(rowValues, "NULL")
			}
		}
		values = append(values, "("+strings.Join(rowValues, ", ")+")")
	}

	// SQLServerで1回のINSERTステートメントで挿入できるレコードの上限は1,000件のため、クエリを分割する
	var queries []string
	for i := 0; i < len(values); i += 1000 {
		end := i + 1000
		if end > len(values) {
			end = len(values)
		}
		queries = append(queries, fmt.Sprintf("INSERT INTO %s VALUES\n%s;\n", tableName, strings.Join(values[i:end], ",\n")))
	}

	// テーブル名.sql として出力
	fileName := fmt.Sprintf("%s.sql", tableName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, query := range queries {
		_, err := writer.WriteString(query)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}
	writer.Flush()

	fmt.Printf("Data successfully exported to %s\n", fileName)
}
