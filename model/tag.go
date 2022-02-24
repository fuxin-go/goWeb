package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Id      int    `gorm:"primaryKey;autoIncrement"`
	TagName string `gorm:"size:100;not null"`
}
