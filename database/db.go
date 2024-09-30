package database

import (
	"log"
	"tsa-innovation-lab/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	DB.AutoMigrate(&models.Contact{})
}
