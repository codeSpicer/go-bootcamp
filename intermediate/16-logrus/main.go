package main

import (
	"github.com/sirupsen/logrus"
)

func main() {

	log := logrus.New()

	// set info level
	log.SetLevel(logrus.InfoLevel)

	// set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// logging examples
	log.Info("THis is an info message")
	log.Warn("this is an warning message")
	log.Error("this is an error message")

	log.WithFields(logrus.Fields{
		"username": "akshat",
		"method":   "GET",
		"endpoint": "/v1/users/profile",
	}).Info("User data profile successfully retrieved")

}
