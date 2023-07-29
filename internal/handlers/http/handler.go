package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.SugaredLogger
}

func NewHandler(logger *zap.SugaredLogger) *Handler {
	return &Handler{logger}
}

func (h Handler) Init(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("/api")
	api.Get("/swagger/*", swagger.HandlerDefault) // default

	api.Get("/ping", h.Ping)

}

// Ping
// @Summary Ping
// @Tags service
// @Description Ping
// @ModuleID Зштп
// @Accept  json
// @Produce  json
//
//	@Success 200 {object} responses.Response
//
// @Failure 400,401,500,503 {null} null
// @Router /ping [get]
func (h Handler) Ping(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"ok":      true,
		"message": "pong",
	})
}
