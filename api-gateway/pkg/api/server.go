package server

import (
	"HireoGateWay/pkg/api/handler"
	"HireoGateWay/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, employerHandler *handler.EmployerHandler, jobSeekerHandler *handler.JobSeekerHandler, jobHandler *handler.JobHandler, chatHandler *handler.ChatHandler, videocallHandler *handler.VideoCallHandler, notihandler *handler.NotificationHandler) *ServerHTTP {

	router := gin.New()

	router.Use(gin.Logger())

	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/exit", videocallHandler.ExitPage)
	router.GET("/error", videocallHandler.ErrorPage)
	router.GET("/index", videocallHandler.IndexedPage)

	// Route for admin auth
	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	// Route for employer auth
	router.POST("/employer/signup", employerHandler.EmployerSignUp)
	router.POST("/employer/login", employerHandler.EmployerLogin)

	router.POST("/job-seeker/signup", jobSeekerHandler.JobSeekerSignUp)
	router.POST("/job-seeker/login", jobSeekerHandler.JobSeekerLogin)

	router.Use(middleware.JobSeekerAuthMiddleware())
	{
		router.GET("/job-seeker/view-jobs", jobHandler.ViewAllJobs)
		router.GET("/job-seeker/jobs", jobHandler.GetJobDetails)

		router.POST("/job-seeker/apply-job", jobHandler.ApplyJob)
		router.GET("/job-seeker/saved-jobs", jobHandler.GetASavedJob)
		router.POST("/job-seeker/save-jobs", jobHandler.SaveAJob)
		router.DELETE("/job-seeker/saved-jobs", jobHandler.DeleteSavedJob)
		router.POST("/job-seeker/apply-saved-job", jobHandler.ApplySavedJob)
		router.GET("/notification", notihandler.GetNotification)
	}

	router.Use(middleware.EmployerAuthMiddleware())
	{
		router.POST("/employer/job-post", jobHandler.PostJobOpening)
		router.GET("/employer/all-job-postings", jobHandler.GetAllJobs)
		router.GET("/employer/job-postings", jobHandler.GetAJob)
		router.DELETE("/employer/job-postings", jobHandler.DeleteAJob)
		router.PUT("/employer/job-postings", jobHandler.UpdateAJob)

		router.GET("/employer/company", employerHandler.GetCompanyDetails)
		router.PUT("/employer/company", employerHandler.UpdateCompany)

		router.GET("/employer/chat", chatHandler.EmployerMessage)
		router.GET("/employer/chats", chatHandler.GetChat)
		router.GET("/group/:groupID/chat", chatHandler.GroupMessage)

		router.GET("/employer/get-applicants", jobHandler.GetApplicants)
		router.POST("/employer/schedule-interview", jobHandler.ScheduleInterview)
		router.GET("/employer/interviews", jobHandler.GetInterviews)
		//router.GET("/employer/interviews", jobHandler.GetAnApplicant)

	}

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :8000")
	err := s.engine.Run(":8000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
