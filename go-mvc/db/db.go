package db

import (
	"fmt"

	"github.com/pixcelo/go-mvc/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
