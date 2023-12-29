package main

import (
	"context"
	"github.com/pworld/loggers"
)

func main() {
	logger := loggers.NewLogger()

	// Create a context with some values
	ctx := context.WithValue(context.Background(), "requestID", "12345")
	ctx = context.WithValue(ctx, "userID", "user-67890")

	// Log with context information
	logger.Info("User accessed resource", "GET", "/resource", 200, ctx.Value("userID"), ctx.Value("requestID"))
}
