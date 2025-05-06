package middleware

import (
    "time"
    "github.com/gin-gonic/gin"
    "go-gin-microservice-advanced/pkg/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        duration := time.Since(start)
        logger.Log.WithFields(map[string]interface{}{
            "status": c.Writer.Status(),
            "method": c.Request.Method,
            "path": c.Request.URL.Path,
            "duration": duration,
        }).Info("request completed")
    }
}