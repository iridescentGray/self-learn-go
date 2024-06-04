package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
打印到控制台/文件 二选一
*/
func main() {
	logFile, _ := os.OpenFile("./app1.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	logrus.SetOutput(logFile)
	// logrus.SetOutput(os.Stderr)
	logrus.Info("o ok ok ok")

}
