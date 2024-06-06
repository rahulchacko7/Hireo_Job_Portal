package middleware

import (
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JobSeekerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logEntry := logging.GetLogger().WithField("context", "JobSeekerAuthMiddleware")
		logEntry.Info("Processing job seeker authentication middleware")

		tokenHeader := c.GetHeader("Authorization") // Note: "Authorization" should be capitalized
		logEntry.Infof("Token Header: %v", tokenHeader)

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
		tokenClaims, err := helper.ValidateTokenJobSeeker(tokenpart)
		if err != nil {
			logEntry.Errorf("Invalid Token: %v", err)
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error()) // Updated error message
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		jobseekerID := int32(tokenClaims.Id)
		c.Set("id", jobseekerID)

		logEntry.Info("Job seeker authenticated successfully")
		c.Next()
	}
}
