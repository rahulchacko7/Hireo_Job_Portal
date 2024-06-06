package handler

import (
	"net/http"

	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	GRPC_Client interfaces.AdminClient
}

func NewAdminHandler(adminClient interfaces.AdminClient) *AdminHandler {
	return &AdminHandler{
		GRPC_Client: adminClient,
	}
}

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "AdminHandler.LoginHandler")
	logEntry.Info("Processing login request")

	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	admin, err := ad.GRPC_Client.AdminLogin(adminDetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during Admin RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Login successful for admin user")
	success := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AdminHandler) AdminSignUp(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "AdminHandler.AdminSignUp")
	logEntry.Info("Processing signup request")

	var adminDetails models.AdminSignUp
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.WithField("email", adminDetails.Email).Info("Creating new admin")

	admin, err := ad.GRPC_Client.AdminSignUp(adminDetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during Admin RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Signup successful for admin user")
	success := response.ClientResponse(http.StatusOK, "Admin created successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}
