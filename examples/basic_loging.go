package main

import "github.com/pworld/loggers"

func main() {
	logger := loggers.NewLogger()

	logger.Info("Simple info message")
	logger.Error("Error occurred", "POST", "/api/error", 500)
	logger.Fatal("Fatal error encountered")
	logger.Warn("Warning: resource limit approaching")
	logger.Debug("Debugging data: ", "GET", "/api/debug", 200)
}
