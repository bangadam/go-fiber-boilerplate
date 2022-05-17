package user_repository

import (
	"context"
	"errors"

	"github.com/bangadam/go-fiber-boilerplate/internal/core/domain"
	"github.com/bangadam/go-fiber-boilerplate/internal/core/port"
	"gorm.io/gorm"
)

type userRepository struct {
	mysql *gorm.DB
}

func NewUserMysql(mysql *gorm.DB) port.UserRepository {
	return &userRepository{
		mysql: mysql,
	}
}

/**
 * GetUserByUsername
 * @param ctx context.Context
 * @param username string
 * @return (*domain.User, error)
 */
func (instance *userRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	err := instance.mysql.Where("user_name = ?", username).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}