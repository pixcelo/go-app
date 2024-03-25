package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func main() {
	filepath := ""
	readExcel(filepath)
}

func readExcel(filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// シート名を取得
	sheets := f.GetSheetMap()

	// シートごとにデータを処理
	for _, sheet := range sheets {
		// シート名が "Sheet1" の場合のみ処理を行う
		targetSheetName := "Sheet1"
		if sheet == targetSheetName {
			// シートからデータを読み込む
			rows, err := f.GetRows(sheet)
			if err != nil {
				return err
			}

			// 1行目がヘッダーなので2行目から処理を開始
			for i, row := range rows {
				if i > 0 {
					bridgeID := row[0]
					noTenken := row[1]
					noKeikan := row[2]
					filePath := row[3]
					fileName := row[4]

					// ファイルをコピーする
					err = copyFile(filePath, fileName, bridgeID, noTenken, noKeikan)
					if err != nil {
						//return err
						//fmt.Println(err)
						fmt.Println(filePath)
					}
				}
			}
		}
	}

	return nil
}

func copyFile(srcPath, fileName, bridgeID, noTenken, noKeikan string) error {
	// 出力先ディレクトリを作成
	dstDir := filepath.Join("Uploads", filepath.Join(bridgeID, "tenken", noTenken, noKeikan, "sonshouzu"))
	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// ファイルをコピー
	//srcFile := filepath.Join(srcPath, fileName)
	dstFile := filepath.Join(dstDir, fileName)
	err = copyFileContents(srcPath, dstFile)
	if err != nil {
		return fmt.Errorf("failed to copy file %s: %v", srcPath, err)
	}

	return nil
}

func copyFileContents(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// ファイルチェック
	_, err = os.Stat(dst)
	if err == nil {
		// ファイルが既に存在する場合は、ログを出力して処理をスキップ
		logFileExist(src, dst)
		return nil
	} else if !os.IsNotExist(err) {
		// その他のエラーの場合は、エラーを返す
		return err
	}
	// ファイルチェック

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func logFileExist(srcFile, dstFile string) {
	fmt.Printf("File %s already exists at %s\n", filepath.Base(srcFile), dstFile)
}
