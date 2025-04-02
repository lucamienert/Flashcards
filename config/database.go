package config

import (
	"log"

	"github.com/lucamienert/flashcards/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *Config) {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Database), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}, &models.SignInInput{}, &models.SignUpInput{}, &models.UserResponse{})
}
