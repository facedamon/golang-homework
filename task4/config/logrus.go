package config

import (
	"fmt"
	"io"
	"os"

	"github.com/facedamon/golang-homework/blog/global"
	"github.com/sirupsen/logrus"
)

type gormWriter struct {
	logger *logrus.Logger
}

func (g *gormWriter) Printf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	g.logger.Debugln(str)
}

func InitLogrus() *logrus.Logger {
	logFile, err := os.OpenFile("blob.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithError(err).Fatal("无法创建日志文件")
		return nil
	}
	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(logFile, os.Stdout))
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		//TimestampFormat:  "2015-09-12 12:44:03",
	})
	global.Logger = logger
	return logger
}
