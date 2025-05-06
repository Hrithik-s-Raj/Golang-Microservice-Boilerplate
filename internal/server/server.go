package server

import (
    "github.com/gin-gonic/gin"
    "go-gin-microservice-advanced/internal/routes"
    "go-gin-microservice-advanced/pkg/config"
    "go-gin-microservice-advanced/pkg/db"
    "go-gin-microservice-advanced/pkg/logger"
    "go-gin-microservice-advanced/pkg/middleware"
)

func Run(env string) error {
    config.LoadEnv(env)
    logger.InitLogger()
    db.ConnectMongo()
    r := gin.New()
    r.Use(gin.Recovery())
    r.Use(middleware.LoggerMiddleware())
    r.Use(middleware.CORSMiddleware())

    routes.RegisterRoutes(r)

    port := config.GetPort()
    return r.Run(":" + port)
}