# Connect Database

基本パッケージの追加
```bash
go mod init github.com/pixcelo/go-db
go get -u gorm.io/gorm
```

ドライバーの追加
```
go get gorm.io/driver/sqlserver
go get github.com/denisenkom/go-mssqldb
```
GORMは公式にMySQL、PostgreSQL、SQLite、SQL Server をサポート


## Local db
server: localhost\SQLEXPRESS
user: testuser
pass: p123

## Reference
- [Gorm](https://github.com/go-gorm)
- [Gorm JP](https://gorm.io/ja_JP/)
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb)