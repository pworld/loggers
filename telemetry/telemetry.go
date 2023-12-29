package telemetry

import (
	"github.com/pworld/loggers/loki"
	"github.com/pworld/loggers/prometheus"
	"os"
)

func InitTelemetry() {
	// Initialize Prometheus
	prometheus.InitPrometheus()

	// Initialize Loki
	loki.NewLokiClient(os.Getenv("LOKI_CLIENT"))
}
