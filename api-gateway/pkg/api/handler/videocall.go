package handler

import (
	logging "HireoGateWay/Logging"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VideoCallHandler struct {
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewVideoCallHandler() *VideoCallHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/Hireo_gateway.log")
	return &VideoCallHandler{
		Logger:  logger,
		LogFile: logFile,
	}
}

func (v *VideoCallHandler) ExitPage(c *gin.Context) {

	v.Logger.Info("Rendering exit page")

	c.HTML(http.StatusOK, "exit.html", nil)
}

func (v *VideoCallHandler) ErrorPage(c *gin.Context) {

	v.Logger.Info("Rendering error page")

	c.HTML(http.StatusOK, "error.html", nil)
}

func (v *VideoCallHandler) IndexedPage(c *gin.Context) {

	v.Logger.Info("Rendering indexed page")

	room := c.DefaultQuery("room", "")
	v.Logger.Infof("Room: %s", room)

	c.HTML(http.StatusOK, "index.html", gin.H{"room": room})
}
