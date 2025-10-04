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
	// Backward-compatible default: same as before
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "study_pal.db"
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = db

	fmt.Println("SQLite database connected:", path)

	// (Optional) create tables on first run:
	// _ = DB.AutoMigrate(&entity.Profile{}, &entity.Course{}, &entity.Task{}, &entity.ProfileCourse{})
}
