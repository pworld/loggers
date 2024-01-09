package loggers

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/pworld/loggers/loki"
	"log"
	"os"
	"time"
)

type Logger struct {
	StandardLogger *log.Logger
	LokiClient     *loki.LokiClient // Assuming you have a LokiClient type
	isLokiActive   bool
	isSentryActive bool
}

// Fields type for structured logging
type Fields map[string]interface{}

// defaultLogger is a package-level Logger instance
var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger()
}

func NewLogger() *Logger {
	return &Logger{
		StandardLogger: log.New(os.Stdout, "", 0),
		LokiClient:     loki.InitializeLokiClient(os.Getenv("LOKI_CLIENT")),
		isLokiActive:   os.Getenv("LOKI_ACTIVE") == "1",
		isSentryActive: os.Getenv("SENTRY_ACTIVE") == "1",
	}
}

// Package-level logging functions
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// genericLog handles the core logging logic for any log level
func (l *Logger) genericLog(level string, args ...interface{}) {
	var message string

	// Determine the message based on the number of arguments
	switch len(args) {
	case 1:
		// If there's only one argument, treat it as a simple message
		message = l.simpleMessage(level, args[0])
	case 2, 3, 4:
		// If there are 2-4 arguments, treat it as a detailed message
		message = l.detailedMessage(level, args...)
	default:
		// If there are no arguments or more than 4, it's considered invalid
		message = fmt.Sprintf("%s: Invalid logging arguments", level)
	}

	// Log the message and send to external systems (if active)
	l.logAndSend(message, sentry.LevelFatal, "fatal")
}

func (l *Logger) CustomFormatLogMessage(level, description string, args ...interface{}) string {
	// Define your custom formatting logic here
	return fmt.Sprintf("[CustomFormat] %s - %s", level, description)
}

func (l *Logger) StructuredLog(level, message string, data map[string]interface{}) {
	// Convert data map to JSON or another structured format
	jsonData, _ := json.Marshal(data)
	logMessage := fmt.Sprintf("%s: %s - %s", level, message, string(jsonData))
	l.StandardLogger.Println(logMessage)
}

// Specific logging methods
func (l *Logger) Info(args ...interface{}) {
	l.genericLog("INFO", args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.genericLog("ERROR", args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.genericLog("FATAL", args...)
	os.Exit(1)
}

func (l *Logger) Warn(args ...interface{}) {
	l.genericLog("WARN", args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.genericLog("DEBUG", args...)
}

// Helper methods for message formatting
func (l *Logger) simpleMessage(level string, arg interface{}) string {
	if msg, ok := arg.(string); ok {
		return fmt.Sprintf("%s: %s", level, msg)
	}
	return fmt.Sprintf("%s: Invalid argument type", level)
}

func (l *Logger) detailedMessage(level string, args ...interface{}) string {
	if len(args) < 2 {
		return fmt.Sprintf("%s: Insufficient arguments for detailed logging", level)
	}
	description, ok := args[0].(string)
	if !ok {
		return fmt.Sprintf("%s: Invalid argument type for description", level)
	}
	method, ok := args[1].(string)
	if !ok {
		return fmt.Sprintf("%s: Invalid argument type for method", level)
	}
	path := "unknown" // default value
	status := 0       // default value
	if len(args) > 2 {
		path, ok = args[2].(string)
		if !ok {
			return fmt.Sprintf("%s: Invalid argument type for path", level)
		}
	}
	if len(args) > 3 {
		status, ok = args[3].(int)
		if !ok {
			return fmt.Sprintf("%s: Invalid argument type for status", level)
		}
	}
	return formatLogMessage(level, description, method, path, status)
}

func formatLogMessage(logType, description, method, path string, status int) string {
	return fmt.Sprintf("[%s] %s %s - %s %s Status: %d",
		time.Now().Local().Format("02-Jan-2006 15:04:05"), logType, description, method, path, status)
}

func (l *Logger) logAndSend(message string, sentryLevel sentry.Level, lokiLevel string) {
	l.StandardLogger.Println(message) // Duplicate logging
	l.sendToSentry(sentryLevel, message)
	l.sendToLoki(lokiLevel, message)
}

func (l *Logger) sendToSentry(level sentry.Level, message string) {
	if l.isSentryActive {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetLevel(level)
			sentry.CaptureMessage(message)
		})
	}
}

func (l *Logger) sendToLoki(level, message string) {
	if l.isLokiActive {
		if err := l.LokiClient.SendLog(level, message); err != nil {
			log.Printf("Failed to send log to Loki: %v\n", err)
		}
	}
}
