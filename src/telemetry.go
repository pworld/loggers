package src

import (
	"os"
)

func InitTelemetry() {
	// Initialize Prometheus
	InitPrometheus()

	// Initialize Loki
	NewLokiClient(os.Getenv("LOKI_CLIENT"))
}
