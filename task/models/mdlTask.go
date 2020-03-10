package models

import "github.com/jinzhu/gorm"

// Task Taskモデル
type Task struct {
	gorm.Model
	Text   string
	Stat   int
	UserID int
}
