package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	config "github.com/utf2/utf-account-service/internal/config/app"
	"github.com/utf2/utf-account-service/internal/logger"
	"github.com/utf2/utf-account-service/internal/server"
)

type App struct {
	logger *slog.Logger
	config *config.Config
}

func New(logger *slog.Logger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (app *App) Run(wg *sync.WaitGroup) error {
	const op = "app.Run()"

	log := app.logger.With(
		slog.String("op", op),
	)

	server, err := server.New(app.config, log)
	if err != nil {
		log.Error("error while creating a new server", logger.Error(err))
		os.Exit(1)
	}

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(fmt.Errorf("error while starting server: %w", err).Error())
			os.Exit(1)
		}
	}()

	log.Debug("server is started", slog.String("port", app.config.HttpServer.Port))
	wg.Done()

	<-interruptChannel
	log.Debug("interrupt signal received, shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error(fmt.Errorf("error while shutting down server: %w", err).Error())
		os.Exit(1)
	}

	log.Debug("server gracefully stopped")
	return nil
}
