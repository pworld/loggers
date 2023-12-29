package main

import (
	"github.com/pworld/loggers"
)

// Usage
func main() {
	logger := loggers.NewLogger()
	message := logger.CustomFormatLogMessage("INFO", "Custom formatted message")
	logger.Info(message)
}
