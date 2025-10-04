package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "study_pal.db" // resolves to /app/study_pal.db in container
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = db

	fmt.Println("SQLite database connected:", path)
}
