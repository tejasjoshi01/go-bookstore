package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:password/books?charset=utf8&parseTime=True") // replace password

	if err != nil {
		fmt.Println("Failed to connect to Database: ", err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
