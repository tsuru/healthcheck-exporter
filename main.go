package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8080", nil)
}
