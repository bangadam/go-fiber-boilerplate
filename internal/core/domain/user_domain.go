package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type User struct {
	ID uint64 `gorm:"primary_key;column:id"`
	CreatedBy uint64 `gorm:"column:created_by"`
	UpdatedBy uint64 `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Name     *string `gorm:"column:name"`
	UserName *string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
}

func (u *User) IsEmpty() bool {
	return u == nil
}
func (u *User) ComparePassword(password string) bool {
	return u.Password == password
}

func (u *User) GenerateTokenAccess() (*string, error) {
	mySigningKey := []byte(viper.GetString("server.secret_key"))
	// Create the Claims
	claims := &jwt.MapClaims{
		"id": u.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}
