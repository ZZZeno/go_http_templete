package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func Logger(filename string) *logrus.Logger {
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	fileName := path.Join(logFilePath, filename)
	writer, _ := rotatelogs.New(
		fileName+".%Y-%m-%d-%H:%M",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(7*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Minute),
	)
	logger := logrus.New()

	logger.Out = writer

	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger("default")
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
