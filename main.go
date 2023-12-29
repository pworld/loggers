package main

import (
	"os"
)

func main() {
	// Initialize environment variables if needed
	os.Setenv("LOKI_ACTIVE", "1")                    // Activate Loki logging
	os.Setenv("SENTRY_ACTIVE", "1")                  // Activate Sentry logging
	os.Setenv("LOKI_CLIENT", "your-loki-client-url") // Set Loki client URL
	os.Setenv("SENTRY_DSN", "sentry-dsn")            // Set Loki client URL
	// Set other required environment variables here...

	// Initialize Sentry (if you have a separate init function in your loggers package)
	// cleanup := loggers.InitSentry()
	// defer cleanup()

	// Perform some logging actions
	InfoLog("Application started", "main", "/", 200)
	DebugLog("This is a debug message", "main", "/debug", 200)

	FatalLog("Fatal error encountered", "main", "/fatal", 500)
}
