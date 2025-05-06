package main

import (
    "log"
    "os"
    "go-gin-microservice-advanced/internal/server"
)

func main() {
    env := os.Getenv("ENV")
    if env == "" {
        env = "local"
    }
    log.Printf("Starting server in %s mode...", env)
    if err := server.Run(env); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}