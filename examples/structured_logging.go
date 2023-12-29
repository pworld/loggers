package main

func main() {
	logger := loggers.NewLogger()

	// Structured logging
	logger.WithFields(loggers.Fields{
		"user_id": 1234,
		"event":   "login",
	}).Info("User login")
}
