package main

import (
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("这是一个普通的日志消息")
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	// Use logrus for standard log output
	// Note that `log` here references stdlib's log
	// Not logrus imported under the name `log`.
	log.SetOutput(logger.Writer())
	log.Println("121231231")
}
