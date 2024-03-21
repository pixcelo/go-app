# インサート用SQLを作成

## 想定するExcel
- 1行目：物理名
- 2行目：論理名
- 3行目～：データ

## パッケージ
```bash
go get -u github.com/xuri/excelize/v2
```

## Usage
```bash
go run main.go {table_name} {file_path}
```

## Reference
- [excelize](https://pkg.go.dev/github.com/xuri/excelize/v2)