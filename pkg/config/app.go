package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "root:vedansh@321@tcp(localhost)/books_management?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("Connected to DB updated")
}

func GetDB() *gorm.DB {
	return db
}
