package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	//ID int `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"size:255;not null"`
	TagId    int    `gorm:"not null"`
	CateId   int    `gorm:"not null"`
	Content  string
	Desc     string
	Status   int
	UserId   int
	Username string
}

type ArticleCate struct {
	Article
	CateName string `gorm:"cate_name" json:"cate_name"`
	TagName  string `gorm:"tag_name" json:"tag_name"`
}
