package main

import (
	"context"
	"log"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/datasdk/pkg/server"
	"github.com/tinygodsdev/places-tg-bot/internal/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.NewStdLogger()
	logger.Info("storage created", "type", cfg.ServerType)

	client, err := server.NewClient(cfg.ServerType, logger)
	if err != nil {
		logger.Fatal("failed to create client", "error", err)
	}

	err = client.HealthCheck(ctx)
	if err != nil {
		logger.Fatal("failed to health check", "error", err)
	}
	logger.Info("health check passed")

}
