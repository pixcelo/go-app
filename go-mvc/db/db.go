package db

import (
	"fmt"

	"github.com/pixcelo/go-mvc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conf := config.Config{
		DbHost:     "service_postgres",
		DbUser:     "postgres",
		DbPassword: "password",
		DbName:     "testdb01",
		DbPort:     "5433"}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DbHost, conf.DbUser, conf.DbPassword, conf.DbName, conf.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
