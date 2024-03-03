package questlister

import (
	"fmt"

	"github.com/Kharec/questlister/internal/models"
	"gorm.io/gorm"
)

type QuestLister struct {
	DB *gorm.DB
}

func QL(db *gorm.DB) *QuestLister {
	return &QuestLister{DB: db}
}

func (ql *QuestLister) Create(title string) error {
	quest := models.Quest{Title: title}
	result := ql.DB.Create(&quest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ql *QuestLister) GetAll() []models.Quest {
	var quests []models.Quest
	ql.DB.Find(&quests)
	return quests
}

func (ql *QuestLister) MarkAsDone(title string) error {
	result := ql.DB.Model(&models.Quest{}).Where("title = ?", title).Update("completed", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("quest not found")
	}
	return nil
}

func (ql *QuestLister) CleanQuests() error {
	result := ql.DB.Where("completed = ?", true).Delete(&models.Quest{})
	return result.Error
}
