package main

import (
	"github.com/cory-johannsen/gomud"
	"log"
)

func main() {
	mud, err := gomud.InitializeEngine()
	if err != nil {
		panic(err)
	}
	log.Printf("Starting server on port %s", mud.Config.Port)
	mud.Server.Start()
}
