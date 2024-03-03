package questlister

import "fmt"

func DisplayHelp() {
	fmt.Println("Manage your todos (quests) in your terminal.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("    ql <command> <subcommand> [flag]")
	fmt.Println()
	fmt.Println("CORE COMMANDS")
	fmt.Println("    add [title]:\t\tAdd a new item with the given name.")
	fmt.Println("    achieve [title]:\tMark the item with the given ID as completed.")
	fmt.Println("    list:\t\tList all items, including the achieved ones.")
	fmt.Println("    clean:\t\tClean the achieved quests from the database.")
	fmt.Println("    help:\t\tThis message is shown if an unknown command is entered.")
}
