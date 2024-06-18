package main

import (
	"log"
	"notification/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("cannot start server:", err)
	} else {
		server.Start()
	}
}
