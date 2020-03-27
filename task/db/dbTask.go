package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // sqlite3
	c "github.com/nishikawa-ssc/todo_test/task/constant"
	m "github.com/nishikawa-ssc/todo_test/task/models"
)

var gd *gorm.DB

// Data データアクセス
func Data() *gorm.DB {
	return gd
}

// Init DB初期化
func Init() {
	gd, _ = gorm.Open("sqlite3", c.DbName)

	// ログを有効にする
	gd.LogMode(true)

	// マイグレート
	gd.AutoMigrate(&m.Task{})
}

// Close DBクローズ
func Close() {
	gd.Close()
}
