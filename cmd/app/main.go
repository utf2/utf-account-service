package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/utf2/utf-account-service/internal/app"
	config "github.com/utf2/utf-account-service/internal/config/app"
	"github.com/utf2/utf-account-service/internal/logger"
)

// @title utf-account-service
// @version 1.0.0
// @description service responsible for operations with students and teachers accounts

// @contact.name Pyankov Daniil
// @contact.url http://t.me/lifelessdev
// @contact.email pyankovdaniildev@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
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

	app := app.New(logger, cfg)

	wg := sync.WaitGroup{}
	wg.Add(1)
	if err := app.Run(&wg); err != nil {
		logger.Error(fmt.Errorf("error while running the app: %w", err).Error())
		os.Exit(1)
	}

	wg.Wait()
	logger.Debug("exiting main() function...")
}
