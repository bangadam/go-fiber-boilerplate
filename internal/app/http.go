package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	Postgres *gorm.DB
	R		 *fiber.App
	Logger *zap.Logger
}

func (h *Handler) SetRoutes() {
	h.R.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	h.R.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Talk is cheap. Show me the code.")
	})

	// Init repository

	// Init service or business logic

	// Init handler
}