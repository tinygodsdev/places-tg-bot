package main

import (
	"context"
	"log"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/datasdk/pkg/server"
	"github.com/tinygodsdev/places-tg-bot/internal/bot"
	"github.com/tinygodsdev/places-tg-bot/internal/config"
	"github.com/tinygodsdev/places-tg-bot/internal/user"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.NewStdLogger()

	client, err := server.NewClient(cfg.ServerType, logger)
	if err != nil {
		logger.Fatal("failed to create client", "error", err)
	}

	err = client.HealthCheck(ctx)
	if err != nil {
		logger.Fatal("failed to health check", "error", err)
	}
	logger.Info("client created", "type", cfg.ServerType, "info", "health check passed")

	userStore, err := user.NewMongoStorage(cfg.Config)
	if err != nil {
		logger.Fatal("failed to create user storage", "error", err)
	}
	logger.Info("user storage created")

	b, err := bot.New(cfg, client, logger, userStore)
	if err != nil {
		logger.Fatal("failed to create bot", "error", err)
	}
	logger.Info("connected to telegram")

	if err := b.Setup(); err != nil {
		logger.Fatal("failed to setup bot", "error", err)
	}
	logger.Info("bot setup completed")

	defer b.Stop()
	b.Start()

}
