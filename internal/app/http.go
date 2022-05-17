package app

import (
	"gorm.io/gorm"

	"github.com/bangadam/go-fiber-boilerplate/internal/adapter/inbound/auth_handler"
	"github.com/bangadam/go-fiber-boilerplate/internal/adapter/outbound/user_repository"
	"github.com/bangadam/go-fiber-boilerplate/internal/core/service/auth_service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handlers struct {
	Mysql      *gorm.DB
	R             *fiber.App
	Logger        *zap.Logger
}

func (h *Handlers) SetupRouter() {
	h.R.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})
	h.R.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Talk Is Cheap Show Me Your Code")
	})
	
	//initialize Repository
	userRepository := user_repository.NewUserMysql(h.Mysql)

	//initialize bussiness
	authService := auth_service.NewAuthService(h.Logger, userRepository)

	//handlers initialize
	auth_handler.NewAuthHandler(h.R, authService)
}
