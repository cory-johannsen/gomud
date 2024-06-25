package main

import (
	"github.com/cory-johannsen/gomud"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Initialize the randomizer
	rand.Seed(time.Now().Unix())

	mud, err := gomud.InitializeEngine()
	if err != nil {
		panic(err)
	}
	log.Printf("Starting server on port %s", mud.Config.Port)
	mud.Server.Start()
}
