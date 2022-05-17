package port

import (
	"context"

	"github.com/bangadam/go-fiber-boilerplate/internal/core/domain"
)

// Repository is the interface for the port layer.
type (
	UserRepository interface {
		GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	}
)