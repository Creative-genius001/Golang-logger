package main

import (
	"github.com/Creative-genius001/go-logger"
)

func main() {
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")

	logger.Error("Something went wrong")
}
