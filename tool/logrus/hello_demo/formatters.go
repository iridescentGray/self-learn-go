package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{}) 默认是TextFormatter
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A group of walrus emerges from the ocean")

	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A group of walrus emerges from the ocean")

}
