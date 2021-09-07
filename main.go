package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/Basketball?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed to Open because", err)
	} else {
		log.Println("Connection Established")
	}

}
