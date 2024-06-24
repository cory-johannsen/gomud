package main

import (
	"github.com/cory-johannsen/gomud/engine"
	"log"
	"math/rand"
	"os"
	"time"
)

const DefaultPort = "7000"

func main() {
	// Initialize the randomizer
	rand.Seed(time.Now().Unix())

	// Extract the server port from the command line
	arguments := os.Args
	var port string
	if len(arguments) == 1 {
		log.Println("No port specified, using default.")
		port = DefaultPort
	} else {
		port = arguments[1]
	}
	log.Printf("Starting server on port %s", port)
	server := &engine.Server{
		Port: port,
	}
	server.Start()
}
