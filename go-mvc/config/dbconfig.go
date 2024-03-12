package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitDB() *gorm.DB {
	config := DBConfig{
		Host:     "127.0.0.1",
		Port:     "1433",
		User:     "testuser",
		Password: "pass123",
		Name:     "testDb",
	}

	// 接続文字列 dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
