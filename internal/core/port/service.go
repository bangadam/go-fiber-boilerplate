package port

import (
	"context"

	"github.com/bangadam/go-fiber-boilerplate/internal/core/domain"
)

// Service is the interface for the port layer.
type (
	AuthService interface {
		Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error)
	}
)