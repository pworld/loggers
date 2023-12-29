package main

import (
	"fmt"
	sentry "github.com/getsentry/sentry-go"
	"github.com/pworld/loggers/src"
	"log"
	"os"
	"time"
)

var (
	LokiClient                   *src.LokiClient
	isLokiActive, isSentryActive bool
)

func init() {
	// Initialize Loki client with values from environment variables
	// LokiClient = src.NewLokiClient(os.Getenv("LOKI_CLIENT"))

	// Set flags for active loggers based on environment variables
	isLokiActive = os.Getenv("LOKI_ACTIVE") == "1"
	isSentryActive = os.Getenv("SENTRY_ACTIVE") == "1"
}

func sendToLoki(level, message string) {
	if isLokiActive {
		if err := LokiClient.SendLog(level, message); err != nil {
			log.Printf("Failed to send log to Loki: %v\n", err)
		}
	}
}

func sendToSentry(level sentry.Level, message string) {
	if isSentryActive {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetLevel(level)
			sentry.CaptureMessage(message)
		})
	}
}

func formatLogMessage(logType, description, method, path string, status int) string {
	return fmt.Sprintf("[%s] %s %s - %s %s Status: %d",
		time.Now().Local().Format("02-Jan-2006 15:04:05"), logType, description, method, path, status)
}

// InfoLog logs an info level message.
func InfoLog(description string, method, path string, status int, args ...interface{}) {
	message := formatLogMessage("INFO", fmt.Sprintf(description, args...), method, path, status)
	log.Println(message)
	sendToSentry(sentry.LevelInfo, message)
	sendToLoki("info", message)
}

// DebugLog logs a debug level message.
func DebugLog(description string, method, path string, status int, args ...interface{}) {
	message := formatLogMessage("DEBUG", fmt.Sprintf(description, args...), method, path, status)
	log.Println(message)
	sendToSentry(sentry.LevelDebug, message)
	sendToLoki("debug", message)
}

// TraceLog logs a trace level message.
func TraceLog(description string, method, path string, status int, args ...interface{}) {
	message := formatLogMessage("TRACE", fmt.Sprintf(description, args...), method, path, status)
	log.Println(message)
	sendToSentry(sentry.LevelDebug, message)
	sendToLoki("trace", message)
}

// ErrorLog logs an error level message.
func ErrorLog(description string, method, path string, status int, args ...interface{}) {
	message := formatLogMessage("ERROR", fmt.Sprintf(description, args...), method, path, status)
	log.Printf("ERROR: %s\n", message)
	sendToSentry(sentry.LevelError, message)
	sendToLoki("error", message)
}

// FatalLog logs a fatal level message and terminates the program.
func FatalLog(description string, method, path string, status int, args ...interface{}) {
	message := formatLogMessage("FATAL", fmt.Sprintf(description, args...), method, path, status)
	sendToSentry(sentry.LevelFatal, message)
	sendToLoki("fatal", message)
	log.Fatalf("FATAL: %s\n", message)
}
