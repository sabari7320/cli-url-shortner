// internal/db/db.go
package db

import (
	"github.com/sabari7320/cli-url-shortner/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("shorty.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.URL{})

	DB = db
}
