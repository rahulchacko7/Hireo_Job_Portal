package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
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

// EmployerLogin godoc
// @Summary Login for employer
// @Description Process the employer login request
// @Tags Employers Authentication
// @Accept json
// @Produce json
// @Param employer body models.EmployerLogin true "Employer Login details"
// @Success 200 {object} response.Response "Employer authenticated successfully"
// @Failure 400 {object} response.Response "Details not in correct format"
// @Failure 500 {object} response.Response "Cannot authenticate employer"
// @Router /employer/login [post]
func (eh *EmployerHandler) EmployerLogin(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "EmployerLogin")
	logEntry.Info("Processing login request")

	var employerDetails models.EmployerLogin
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Info("Request body bound successfully")
	employer, err := eh.GRPC_Client.EmployerLogin(employerDetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during Employer RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate employer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Login successful for user")
	success := response.ClientResponse(http.StatusOK, "Employer authenticated successfully", employer, nil)
	c.JSON(http.StatusOK, success)
}

// EmployerSignUp godoc
// @Summary Sign up a new employer
// @Description Process the employer signup request
// @Tags Employers Authentication
// @Accept json
// @Produce json
// @Param employer body models.EmployerSignUp true "Employer SignUp details"
// @Success 200 {object} response.Response "Employer created successfully"
// @Failure 400 {object} response.Response "Details not in correct format"
// @Failure 500 {object} response.Response "Cannot create employer"
// @Router /employer/signup [post]
func (eh *EmployerHandler) EmployerSignUp(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "EmployerSignUp")
	logEntry.Info("Processing signup request")

	var employerDetails models.EmployerSignUp
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Info("Request body bound successfully")
	employer, err := eh.GRPC_Client.EmployerSignUp(employerDetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during Employer RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create employer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Signup successful for user")
	success := response.ClientResponse(http.StatusOK, "Employer created successfully", employer, nil)
	c.JSON(http.StatusOK, success)
}

// GetCompanyDetails godoc
// @Summary Fetch company details
// @Description Retrieve details of the company deails with the authenticated employer
// @Tags Employers
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.Response "Company details retrieved successfully"
// @Failure 400 {object} response.Response "Invalid employer ID type"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Failed to fetch company details"
// @Router /employer/company [get]
func (eh *EmployerHandler) GetCompanyDetails(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "GetCompanyDetails")
	logEntry.Info("Fetching company details")

	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Error("Employer ID not found in context")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Infof("Employer ID: %v", employerID)

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Info("Requesting company details from GRPC client")

	companyDetails, err := eh.GRPC_Client.GetCompanyDetails(employerIDInt)
	if err != nil {
		logEntry.WithError(err).Error("Error during GetCompanyDetails RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to fetch company details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Company details retrieved successfully")
	response := response.ClientResponse(http.StatusOK, "Company details retrieved successfully", companyDetails, nil)
	c.JSON(http.StatusOK, response)
}

// UpdateCompany godoc
// @Summary Update company details
// @Description Update details of the company associated with the authenticated employer
// @Tags Employers
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param employerDetails body models.EmployerDetails true "Employer details to update"
// @Success 200 {object} response.Response "Company updated successfully"
// @Failure 400 {object} response.Response "Invalid employer ID type"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Failed to update company"
// @Router /employer/company [put]
func (eh *EmployerHandler) UpdateCompany(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "UpdateCompany")
	logEntry.Info("Processing update company request")

	var employerDetails models.EmployerDetails
	if err := c.ShouldBindJSON(&employerDetails); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Info("Request body bound successfully")
	employerID, ok := c.Get("id")
	if !ok {
		logEntry.Error("Employer ID not found in context")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	employerIDInt, ok := employerID.(int32)
	if !ok {
		logEntry.Error("Invalid employer ID type")
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid employer ID type", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	logEntry.Info("Requesting update company details from GRPC client")
	updatedCompany, err := eh.GRPC_Client.UpdateCompany(employerIDInt, employerDetails)
	if err != nil {
		logEntry.WithError(err).Error("Error during UpdateCompany RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to update company", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	logEntry.Info("Company updated successfully")
	response := response.ClientResponse(http.StatusOK, "Company updated successfully", updatedCompany, nil)
	c.JSON(http.StatusOK, response)
}
