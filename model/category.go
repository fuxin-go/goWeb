package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Id       int    `gorm:"primaryKey;"`
	CateName string `gorm:"not null;unique"`
}
