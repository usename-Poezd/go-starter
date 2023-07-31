package app

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/usename-Poezd/go-starter/internal/config"
	"github.com/usename-Poezd/go-starter/internal/handlers/http"
	"github.com/usename-Poezd/go-starter/pkg/logger"
)

// @title API
// @version 1.0
// @description API
// @host localhost:8000
// @BasePath /api
// Run initializes whole application.
func Run() {
	log := logger.Get()
	_, err := config.Init()
	if err != nil {
		log.Info().Err(err).Msg("cannot load config")
	}

	app := fiber.New()
	logger := logger.NewConsole()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
        Logger: &logger,
    }))

	h := http.NewHandler()
	h.Init(app)

	app.Listen(":8000")
}
