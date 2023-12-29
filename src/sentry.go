package src

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func InitSentry() func() {
	sentryDsn := os.Getenv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn, // Use the DSN directly as a string
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Named cleanup function
	sentryCleanup := func() {
		sentry.Flush(2 * time.Second)
	}

	return sentryCleanup
}
