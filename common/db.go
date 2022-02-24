package common

import (
	"fmt"
	"gin-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := Conf.DB.DriverName
	host := Conf.DB.Host
	port := Conf.DB.Port
	username := Conf.DB.Username
	password := Conf.DB.Password
	database := Conf.DB.Database
	charset := Conf.DB.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	db.LogMode(true)
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Category{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
