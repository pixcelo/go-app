package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBConfig struct {
	Database Database `yaml:database`
}

type Database struct {
	Host     string `yaml:host`
	Port     string `yaml:port`
	User     string `yaml:user`
	Password string `yaml:password`
	Name     string `yaml:name`
}

func InitDB() (*gorm.DB, error) {
	viper.SetConfigName("config")  // 設定ファイルの名前
	viper.SetConfigType("yml")     // 設定ファイルの形式
	viper.AddConfigPath("config/") // 設定ファイルのパス

	err := viper.ReadInConfig() // 設定ファイルの読み取り
	if err != nil {
		return nil, fmt.Errorf("Error: Reading config.yml file %s \n", err)
	}

	var config DBConfig

	err = viper.Unmarshal(&config) // シリアル化されたデータ構造をstructに変換
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}

	// 接続文字列 dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, err
}
