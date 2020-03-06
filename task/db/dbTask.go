package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // sqlite3
	c "github.com/nishikawa-ssc/todo_test/task/constant"
	m "github.com/nishikawa-ssc/todo_test/task/models"
)

var db *gorm.DB

// Task データアクセス
func Task() *gorm.DB {
	return db
}

// Init DB初期化
func Init() {
	db, _ = gorm.Open("sqlite3", c.DbName)

	// ログを有効にする
	db.LogMode(true)

	// マイグレート
	db.AutoMigrate(&m.Task{})
}

// Close DBクローズ
func Close() {
	db.Close()
}
