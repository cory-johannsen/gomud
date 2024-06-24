package main

import (
	"fmt"
	"github.com/cory-johannsen/gomud/cli"
	"github.com/openengineer/go-repl"
	"log"
)

func main() {
	fmt.Println("Welcome, type \"help\" for more info")

	h := &cli.CommandHandler{}
	h.R = repl.NewRepl(h)

	// start the terminal loop
	if err := h.R.Loop(); err != nil {
		log.Fatal(err)
	}
}
