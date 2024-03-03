package main

import (
	"os"

	"github.com/Kharec/questlister/internal/questlister"
)

func main() {
	questlister.Run(os.Args[1:])
}
