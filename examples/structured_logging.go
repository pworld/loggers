package main

import "github.com/pworld/loggers"

func main() {
	logger := loggers.NewLogger()

	// Structured logging
	logger.StructuredLog("INFO", "User login", loggers.Fields{
		"user_id": 1234,
		"event":   "login",
	})
}
