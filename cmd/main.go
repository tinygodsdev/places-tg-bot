package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
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

	client, err := server.NewClient(cfg.ServerType, logger)
	if err != nil {
		logger.Fatal("failed to create client", "error", err)
	}

	err = client.HealthCheck(ctx)
	if err != nil {
		logger.Fatal("failed to health check", "error", err)
	}
	logger.Info("health check passed")
	logger.Info("client created", "type", cfg.ServerType)

	for {
		points, err := client.GetPoints(ctx,
			data.Filter{},
			data.Group{
				TagLabels: []string{"city"},
			},
		)
		if err != nil {
			logger.Error("failed to get points", "error", err)
		}

		fmt.Printf("points: %+v\n", points)

		time.Sleep(120 * time.Minute)
	}
}
