package middleware

import (
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/logging"
	"HireoGateWay/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logEntry := logging.GetLogger().WithField("middleware", "AdminAuthMiddleware")

		tokenHeader := c.GetHeader("authorization")
		logEntry.Infof("Received token header: %s", tokenHeader)

		if tokenHeader == "" {
			logEntry.Warn("No auth header provided")
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			logEntry.Warn("Invalid Token Format")
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateToken(tokenpart)
		if err != nil {
			logEntry.Errorf("Invalid Token: %s", err.Error())
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		logEntry.Info("Token validated successfully")
		c.Set("tokenClaims", tokenClaims)

		c.Next()
	}
}
