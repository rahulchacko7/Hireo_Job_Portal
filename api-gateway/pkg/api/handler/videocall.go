package handler

import (
	"HireoGateWay/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoCallHandler struct{}

func NewVideoCallHandler() *VideoCallHandler {
	return &VideoCallHandler{}
}

func (v *VideoCallHandler) ExitPage(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "ExitPage")
	logEntry.Info("Rendering exit page")

	c.HTML(http.StatusOK, "exit.html", nil)
}

func (v *VideoCallHandler) ErrorPage(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "ErrorPage")
	logEntry.Info("Rendering error page")

	c.HTML(http.StatusOK, "error.html", nil)
}

func (v *VideoCallHandler) IndexedPage(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "IndexedPage")
	logEntry.Info("Rendering indexed page")

	room := c.DefaultQuery("room", "")
	logEntry.Infof("Room: %s", room)

	c.HTML(http.StatusOK, "index.html", gin.H{"room": room})
}
