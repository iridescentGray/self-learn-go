package main

import (
	log "github.com/sirupsen/logrus"
)

func logInFunc() {
	log.Info("will c")
}

func main() {
	log.SetReportCaller(true)
	logInFunc()
}
