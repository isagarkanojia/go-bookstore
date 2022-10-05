package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect(){

 
	dbURL := "postgres://sagar:postgres@localhost:5432/bookstore"

	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})


	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}