package app

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	_ "github.com/usename-Poezd/go-starter/docs"
	"github.com/usename-Poezd/go-starter/internal/config"
	"github.com/usename-Poezd/go-starter/internal/handlers/http"
	"github.com/usename-Poezd/go-starter/internal/services"
	"github.com/usename-Poezd/go-starter/pkg/logger"
)

// @title API
// @version 1.0
// @description API
// @host localhost:8000
// @BasePath /api
// Run initializes whole application.
func Run() {
	logger := logger.NewConsole()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	_, err := config.Init()
	if err != nil {
		logger.Info().Err(err).Msg("cannot load config")
	}

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	s := services.NewServices()
	h := http.NewHandler(s)
	h.Init(app)

	go func() {
		err := app.Listen(":8000")
		if err != nil {
			logger.Error().Err(err).Msg("Can not start http server")
			stop()
		}
	}()

	<-ctx.Done()
	logger.Info().Msg("--- Shutdonwing service ---")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		logger.Fatal().Err(err).Msg("Unable to shutdown")
	}

	logger.Info().Msg("--- Service is shutdown ---")
}
