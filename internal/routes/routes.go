package routes

import (
    "github.com/gin-gonic/gin"
    "go-gin-microservice-advanced/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
    {
        api.GET("/ping", controllers.Ping)
    }
}