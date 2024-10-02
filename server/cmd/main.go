package main

import (
	"log"
)

func main() {
	mud, err := InitializeEngine()
	if err != nil {
		panic(err)
	}
	log.Printf("Starting server on port %s", mud.Config.Port)
	mud.Server.Start()
}
