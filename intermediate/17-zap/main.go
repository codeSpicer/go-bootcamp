package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {

	// https://pkg.go.dev/go.uber.org/zap#section-readme

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error setting up logger", err)
	}
	defer logger.Sync()

	logger.Info("Logger set up")

	logger.Info("User is logged in", zap.String("username", "akshat"), zap.String("method", "GET"), zap.Int("number", 2))

	// sugar := logger.Sugar()

	// url := "google.com"
	// sugar.Infow("failed to fetch url", "url", url, "attempt", 3, "backoff", time.Second)

	// sugar.Infof("Failed to fetch URL: %s", url)

}
