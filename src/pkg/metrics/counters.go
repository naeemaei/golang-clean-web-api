package metrics

import "github.com/prometheus/client_golang/prometheus"

var DbCalls = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "db_call_total",
		Help: "Number of database calls",
	},
	[]string{"type_name", "operation_name", "status"},
)
