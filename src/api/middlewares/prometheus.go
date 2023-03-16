package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		timer := prometheus.NewTimer(metrics.HttpDuration.WithLabelValues(path, method))

		c.Next() 
		
		timer.ObserveDuration()
	}
}
