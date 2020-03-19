package models

// Task Taskモデル
type Task struct {
	ID     int    `gorm:"primary_key;not null"       json:"id"`
	Text   string `gorm:"type:varchar(200);not null" json:"text"`
	Stat   int    `gorm:"not null"                   json:"stat"`
	UserID int
}
