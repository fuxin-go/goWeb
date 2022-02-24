package model

import "time"

// User
type User struct {
	ID       int    `gorm:"primaryKey"` //设置为主键
	Username string `gorm:"not null;unique;size:255"`
	Password string `gorm:"not null"`
	Mobile   string `gorm:"unique;not null;"`
	CreateAt time.Time
	Status   int
}
