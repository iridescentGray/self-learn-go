package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("A walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
