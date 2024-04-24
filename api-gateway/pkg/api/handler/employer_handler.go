package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployerHandler struct {
	GRPC_Client interfaces.EmployerClient
}

func NewEmployerHandler(employerClient interfaces.EmployerClient) *EmployerHandler {
	return &EmployerHandler{
		GRPC_Client: employerClient,
	}
}

func (eh *EmployerHandler) EmployerLogin(c *gin.Context) {
	var employerDetails models.EmployerLogin
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employer, err := eh.GRPC_Client.EmployerLogin(employerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate employer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Employer authenticated successfully", employer, nil)
	c.JSON(http.StatusOK, success)
}

func (eh *EmployerHandler) EmployerSignUp(c *gin.Context) {
	var employerDetails models.EmployerSignUp

	fmt.Println("gateway", employerDetails.Contact_email)

	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employer, err := eh.GRPC_Client.EmployerSignUp(employerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create employer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Employer created successfully", employer, nil)
	c.JSON(http.StatusOK, success)
}

func (eh *EmployerHandler) PostJobOpening(c *gin.Context) {

	var jobOpening models.JobOpening

	if err := c.ShouldBindJSON(&jobOpening); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	

}
