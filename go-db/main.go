package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	// github.com/denisenkom/go-mssqldb
	//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	dsn := "sqlserver://testuser:p123@localhost:1433?database=testDb"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
