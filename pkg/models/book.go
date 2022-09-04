package models

import (
	"github.com/isagarkanojia/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
type Book struct{
	gorm.Model
	Name string `gorm:"json:name"`
	Author string `gorm:"json:author"`
	Publication string `gorm:"json:publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}