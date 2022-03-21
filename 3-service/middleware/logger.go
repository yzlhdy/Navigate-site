package middleware

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// logger struct
func Loggers() gin.HandlerFunc {
	// Set to JSON format
	filePathText := "log/log"
	linkName := "latest_log.log"
	src, err := os.OpenFile(filePathText, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := logrus.New()
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logf, _ := rotatelogs.New(
		filePathText+"%Y%m%d%H%M.log",

		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(linkName),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logf,
		logrus.FatalLevel: logf,
		logrus.DebugLevel: logf,
		logrus.WarnLevel:  logf,
		logrus.ErrorLevel: logf,
		logrus.PanicLevel: logf,
	}
	// logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{}))
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logger.Hooks.Add(lfHook)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIpText := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		methodText := c.Request.Method
		path := c.Request.URL.Path
		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostName,
			"status":     statusCode,
			"path":       path,
			"method":     methodText,
			"client_ip":  clientIpText,
			"user_agent": userAgent,
			"data_size":  dataSize,
			"spend_time": spendTime,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}

}
