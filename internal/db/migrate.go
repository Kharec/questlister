package db

import (
	"github.com/Kharec/questlister/internal/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(database *gorm.DB) {
	err := database.AutoMigrate(&models.Quest{})
	if err != nil {
		log.Fatal("database migration failed")
	}
}
