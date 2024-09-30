package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	config "github.com/utf2/utf-account-service/internal/config/app"
	"github.com/utf2/utf-account-service/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("error while loading config: %w", err))
		os.Exit(1)
	}

	logger := logger.New(cfg.Env).With(
		slog.String("op", "main()"),
	)

	logger.Debug("logger created successfully", slog.Any("env", cfg.Env))
}
