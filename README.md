# Go Gin Microservice Boilerplate (Advanced)

## Features
- Gin Framework
- MongoDB with environment-based config
- Structured Logging with Logrus
- Graceful CORS & Error Middleware
- Multi-environment support: local, dev, qa, prod

## Getting Started
```bash
cp .env.local .env
go run ./cmd/main.go
```

Set `ENV=production` or others in your shell to switch environments.