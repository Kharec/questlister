package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
