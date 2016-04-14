package main

import (
	log "github.com/Sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Info("A walrus appears")
}
