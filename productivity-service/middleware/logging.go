package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)

		c.Next()

		duration := time.Since(start).Seconds() * 1000
		log := time.Now().UTC().Format(time.RFC3339) +
			" [RequestID: " + requestID + "] " +
			c.Request.Method + " " + c.Request.URL.Path + " - " +
			strconv.Itoa(c.Writer.Status()) + " - Duration: " +
			fmt.Sprintf("%.3fms", duration)
		fmt.Println(log)
	}
}
