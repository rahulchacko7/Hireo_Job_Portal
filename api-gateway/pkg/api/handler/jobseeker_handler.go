package handler

import (
	logging "HireoGateWay/Logging"
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type JobSeekerHandler struct {
	GRPC_Client interfaces.JobSeekerClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewJobSeekerHandler(jobSeekerClient interfaces.JobSeekerClient) *JobSeekerHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/Hireo_gateway.log")
	return &JobSeekerHandler{
		GRPC_Client: jobSeekerClient,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// JobSeekerLogin godoc
// @Summary Job seeker login
// @Description Authenticate job seeker
// @Tags Job Seekers Authentication
// @Accept json
// @Produce json
// @Param request body models.JobSeekerLogin true "Job seeker credentials"
// @Success 200 {object} response.Response "Job seeker authenticated successfully"
// @Failure 400 {object} response.Response "Details not in correct format"
// @Failure 500 {object} response.Response "Cannot authenticate job seeker"
// @Router /job-seeker/login [post]
func (jh *JobSeekerHandler) JobSeekerLogin(c *gin.Context) {

	jh.Logger.Info("Processing job seeker login request")

	var jobSeekerDetails models.JobSeekerLogin
	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		jh.Logger.Error("Details not in correct format")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobSeeker, err := jh.GRPC_Client.JobSeekerLogin(jobSeekerDetails)
	if err != nil {
		jh.Logger.Errorf("Cannot authenticate job seeker: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate job seeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	jh.Logger.Info("Job seeker authenticated successfully")
	success := response.ClientResponse(http.StatusOK, "Job seeker authenticated successfully", jobSeeker, nil)
	c.JSON(http.StatusOK, success)
}

// JobSeekerSignUp godoc
// @Summary Job seeker sign up
// @Description Create a new job seeker account
// @Tags Job Seekers Authentication
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body models.JobSeekerSignUp true "Job seeker details"
// @Success 200 {object} response.Response "Job seeker created successfully"
// @Failure 400 {object} response.Response "Details not in correct format"
// @Failure 401 {object} response.Response "No auth header provided"
// @Failure 500 {object} response.Response "Cannot create job seeker"
// @Router /job-seeker/signup [post]
func (jh *JobSeekerHandler) JobSeekerSignUp(c *gin.Context) {

	jh.Logger.Info("Processing job seeker sign up request")

	var jobSeekerDetails models.JobSeekerSignUp
	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		jh.Logger.Error("Details not in correct format")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobSeeker, err := jh.GRPC_Client.JobSeekerSignUp(jobSeekerDetails)
	if err != nil {
		jh.Logger.Errorf("Cannot create job seeker: %v", err)
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create job seeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	jh.Logger.Info("Job seeker created successfully")
	success := response.ClientResponse(http.StatusOK, "Job seeker created successfully", jobSeeker, nil)
	c.JSON(http.StatusOK, success)
}
