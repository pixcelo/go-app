package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqkx.DB, error) {
	// 環境変数からDB情報を取得
	dbUser := os.Getenv("DB_USERR")          // login-user
	dbPassword := os.Getenv(("DB_PASSWORD")) // login-pass
	dbHost := os.Getenv("DB_HOST")           // db
	dbName := os.Getenv("DB_NAME")           // login-db
	dbPort := os.Getenv("DB_PORT")           // 3306
	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", src)
	if err != nil {
		return nil, fmt.Errorf("failed to Open DB: %w, err")
	}

	// pingで接続の確認
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to Ping DB: %w", err)
	}

	xdb := sqlx.NewDb(db, "mysql")
	return xdb, nil
}
