package metrics

import "github.com/prometheus/client_golang/prometheus"

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of requests.",
	},
	[]string{"path", "method", "status_code"},
)
