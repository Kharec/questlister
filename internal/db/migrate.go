package db

import (
	"github.com/Kharec/questlister/internal/models"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(&models.Quest{})
}
