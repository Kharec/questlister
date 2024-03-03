package models

import "gorm.io/gorm"

type Quest struct {
	gorm.Model
	ID        int    `gorm:"size:3;primary_key;auto_increment:true"`
	Title     string `gorm:"size:256"`
	Completed bool   `gorm:"default:false"`
}
