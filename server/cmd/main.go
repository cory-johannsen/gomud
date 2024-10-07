package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: false,
		DisableColors: false,
	})
	log.SetReportCaller(true)
	mud, err := InitializeEngine()
	if err != nil {
		panic(err)
	}
	log.Printf("Starting server on port %s", mud.Config.Port)
	mud.Server.Start()
}
