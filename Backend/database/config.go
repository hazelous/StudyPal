package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:210306@tcp(127.0.0.1:3306)/studypal?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't Connect to Database")
	} else {
		fmt.Println("Database Connected")
	}
}
