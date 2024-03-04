package questlister

import (
	"os"
	"path/filepath"

	"github.com/Kharec/questlister/internal/db"
)

func Run(args []string) {

	dbPath := filepath.Join(os.Getenv("HOME"), ".config/questlister/ql.db")

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
