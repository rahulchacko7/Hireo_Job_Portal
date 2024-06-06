package middleware

import (
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EmployerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logEntry := logging.GetLogger().WithField("middleware", "EmployerAuthMiddleware")
		logEntry.Info("Processing employer authentication middleware")

		tokenHeader := c.GetHeader("authorization")
		logEntry.Infof("Token header: %s", tokenHeader)

		if tokenHeader == "" {
			logEntry.Error("No auth header provided")
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			logEntry.Error("Invalid Token Format")
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateTokenEmployer(tokenpart)
		if err != nil {
			logEntry.Errorf("Invalid Token: %v", err)
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		employerID := int32(tokenClaims.Id)
		logEntry.Infof("Employer ID: %v", employerID)

		c.Set("id", employerID)

		logEntry.Info("Employer authentication successful")
		c.Next()
	}
}
