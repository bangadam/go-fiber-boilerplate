package auth_service

import (
	"context"
	"errors"
	"net/http"

	"github.com/bangadam/go-fiber-boilerplate/internal/core/domain"
	"github.com/bangadam/go-fiber-boilerplate/internal/core/port"
	"github.com/bangadam/go-fiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrCredentialInvalid = errors.New("credential invalid")
	ErrLoginFailed = errors.New("login failed")
)

type authService struct {
	log *zap.Logger
	userRepository port.UserRepository
}

func NewAuthService(log *zap.Logger, userRepository port.UserRepository) port.AuthService {
	return &authService{
		log:            log,
		userRepository: userRepository,
	}
}

func (instance *authService) Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := instance.userRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		// log error
		instance.log.Error("failed to get user", zap.Error(err))
		return nil, response.New(fiber.StatusInternalServerError, response.WithMessage(ErrLoginFailed.Error()))
	}

	if user.IsEmpty() {
		return nil, response.New(fiber.StatusNotFound, response.WithMessage(ErrUserNotFound.Error()))
	}

	if !user.ComparePassword(req.MD5Password()) {
		return nil, response.New(http.StatusUnauthorized, response.WithMessage(ErrCredentialInvalid.Error()))
	}

	token, err := user.GenerateTokenAccess()
	if err != nil {
		// log error
		instance.log.Error("failed to generate token", zap.Error(err))
		return nil, response.New(http.StatusInternalServerError, response.WithMessage(ErrLoginFailed.Error()))
	}

	return &domain.LoginResponse{
		AccessToken: *token,
	}, nil
}