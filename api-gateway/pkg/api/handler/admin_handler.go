package handler

import (
	"net/http"
	"os"

	"HireoGateWay/Logging"
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewAdminHandler(adminClient interfaces.AdminClient) *AdminHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/Hireo_gateway.log")
	return &AdminHandler{
		GRPC_Client: adminClient,
		Logger:      logger,
		LogFile:     logFile,
	}
}

func (ad *AdminHandler) LoginHandler(c *gin.Context) {

	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AdminHandler) AdminSignUp(c *gin.Context) {

	var adminDetails models.AdminSignUp
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Admin created successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}
