package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)


func Connect(){

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost:5432", "sagar", "bookstore", "postgres") //Build connection string
	fmt.Println(dbUri)

	d, err := gorm.Open("postgres", dbUri)
	
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}