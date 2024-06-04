package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

/*
能打印到控制台,也能打引到文件
*/
func main() {
	writers := []io.Writer{os.Stderr}
	logFile, _ := os.OpenFile("./app2.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	writers = append(writers, logFile)
	logger := logrus.Logger{
		Out: io.MultiWriter(writers...),
		Formatter: &logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Level: logrus.InfoLevel,
	}
	logger.Info("o ok ok ok")
}
