package models

import "time"

// User Userモデル
type User struct {
	ID        int
	Name      string
	PassWord  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
