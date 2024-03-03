package questlister

import (
	"github.com/Kharec/questlister/internal/db"
)

const dbPath = "~/.config/questlister/ql.db"

func Run(args []string) {

	database := db.OpenDatabase(dbPath)
	db.Migrate(database)
	ql := QL(database)

	if len(args) < 1 {
		DisplayHelp()
		return
	}

	switch args[0] {
	case "add":
		ql.Create(args[1])

	case "achieve":
		ql.MarkAsDone(args[1])

	case "list":
		PrintQuests(ql.GetAll())

	case "clean":
		ql.CleanQuests()

	case "help":
		DisplayHelp()

	case "default":
		return
	}
}
