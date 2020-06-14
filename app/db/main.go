package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(db:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db connection error")
	}
	return db
}
