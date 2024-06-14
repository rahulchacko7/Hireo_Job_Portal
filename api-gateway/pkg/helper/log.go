package helper

import (
	logging "HireoGateWay/Logging"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	logrusLogger, logrusLogFile := logging.InitLogrusLogger("./Logging/hireo_jobs_gateway.log")
	defer logrusLogFile.Close()
	return logrusLogger
}
