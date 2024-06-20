package main

import (
	"notification/pkg/config"
	"notification/pkg/di"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		return
	}
	server, diErr := di.InitializeApi(config)
	if diErr != nil {

	} else {
		server.Start()
	}

}
