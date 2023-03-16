package metrics

import "github.com/prometheus/client_golang/prometheus"


var HttpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time",
		Help: "Duration of HTTP requests",
}, []string{"path","method","status_code"})