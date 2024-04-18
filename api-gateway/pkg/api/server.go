package server

import (
	"HireoGateWay/pkg/api/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, employerHandler *handler.EmployerHandler) *ServerHTTP {

	router := gin.New()

	router.Use(gin.Logger())

	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)
	router.POST("/employer/signup", employerHandler.EmployerSignUp)
	router.POST("/employer/login", employerHandler.EmployerLogin)

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :8000")
	err := s.engine.Run(":8000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
