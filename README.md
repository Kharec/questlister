# questlister

A simple todo management command-line tool backed by SQLite.

Don't think of your todo as simple tasks, they are quests and you're a knight.

## Build

```
➜  make
➜  make install # push the binary to ~/bin folder
```

## Usage

```
Manage your todos (quests) in your terminal.

USAGE
    ql <command> <subcommand> [flag]

CORE COMMANDS
    add [title]:		Add a new item with the given name.
    achieve [title]:	Mark the item with the given ID as completed.
    list:		List all items, including the achieved ones.
    clean:		Clean the achieved quests from the database.
    help:		This message is shown if an unknown command is entered.
```

## Database

A SQLite database is backing this tool, it's stored in `~/.config/questlister/ql.db`.

## Example

```
➜  ql list
# nothing yet
```

Adding quests :

```
➜  ql add "write readme"
➜  ql add "clean code"
```

```
➜  ql list              
[] write readme
[] clean code
```

Achieve one quest :

```
➜  ql achieve "write readme"
➜  ql list                  
[X] write readme
[] clean code
```

Delete achieved quests :

```
➜  ql clean
➜  ql list 
[] clean code
```

## License

MIT Licence (a copy is in this repository)
