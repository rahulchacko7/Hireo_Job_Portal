package main

import (
	_ "HireoGateWay/cmd/docs"
	"HireoGateWay/Logging"
	"HireoGateWay/pkg/config"
	"HireoGateWay/pkg/di"
	"log"
)

// @title Go + Gin E-Commerce API
// @title Hireo Jobs API
// @version 1.0
// @description Hire_Jobs is a platform to find your dream job.
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /

func main() {

	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/Hireo_gateway.log")
	defer logrusLogFile.Close()

	config, configErr := config.LoadConfig()
	if configErr != nil {
		logrusLogger.Error("Cannot load config", configErr)
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		logrusLogger.Info("Server Started...")
		server.Start()
	}

}
