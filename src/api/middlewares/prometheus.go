package middlewares

import (
	"GOLANG_CLEAN_WEB_API/src/pkg/metrics"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := float64(time.Since(start) / time.Millisecond)

		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path // مسیر واقعی حتی برای 404
		}

		method := c.Request.Method
		status := c.Writer.Status()

		fmt.Printf("📊 Prometheus Log: path=%q method=%s status=%d duration=%.2fms\n", path, method, status, duration)

		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).
			Observe(duration)
	}
}
