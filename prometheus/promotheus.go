package prometheus

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		port := ":9090" // Change this to the desired port
		log.Println("Serving Prometheus metrics on", port)
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalf("Error starting Prometheus HTTP server: %v", err)
		}
	}()
}
