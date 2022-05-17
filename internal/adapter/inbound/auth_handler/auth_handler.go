package auth_handler

import (
	"github.com/bangadam/go-fiber-boilerplate/internal/core/domain"
	"github.com/bangadam/go-fiber-boilerplate/internal/core/port"
	"github.com/bangadam/go-fiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	app *fiber.App
	authService port.AuthService
}

func NewAuthHandler(app *fiber.App, authService port.AuthService)  {
	authHandler := &authHandler{
		app: app,
		authService: authService,
	}
	
	api := authHandler.app.Group("/api/v1")
	api.Post("/login", authHandler.DoLogin)
}

func (instance *authHandler) DoLogin(c *fiber.Ctx) error {
	req := new(domain.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))
	}

	if err := req.LoginValidation(); err != nil {
		return response.Response(c, response.New(fiber.StatusUnprocessableEntity, response.WithMeta(err)))
	}

	result, err := instance.authService.Login(c.Context(), *req)
	if err != nil {
		return response.Response(c, err)
	}

	return response.Success(c, fiber.StatusOK, response.SuccessMeta(result))
}