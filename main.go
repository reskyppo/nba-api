package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// init Global Variable
var db *gorm.DB
var err error

// Struct for Teams data details
type Teams struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

func main() {
	// Define DB name, DB username, DB password
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/Basketball?charset=utf8&parseTime=True")

	// Check if it's connect to DB or not
	if err != nil {
		log.Println("Connection Failed to Open because", err)
	} else {
		log.Println("Connection Established")
	}
	
	// Create Table Teams 
	db.AutoMigrate(&Teams{})
}
