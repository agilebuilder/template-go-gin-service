// Package middleware provides generic Gin middleware.
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger logs the method, path, status code and latency of every request.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("%s %s %d %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}
