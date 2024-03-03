package questlister

import (
	"fmt"

	"github.com/Kharec/questlister/internal/models"
)

func PrintQuests(quests []models.Quest) {

	if len(quests) == 0 {
		return
	}

	for _, quest := range quests {
		if quest.Completed {
			fmt.Print("[X] ")
		} else {
			fmt.Print("[ ] ")
		}
		fmt.Printf("%s\n", quest.Title)
	}
}
