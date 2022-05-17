package middleware

import (
	"context"
	"errors"
	"strings"

	responseErr "github.com/bangadam/go-fiber-boilerplate/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var (
	ErrMalformedJWT = errors.New("karedensial tidak dapat diketahui")
	ErrUnauthorized = errors.New("kamu tidak memiliki akses")
	ErrDefaultError = errors.New("server sedang sibuk coba beberapa saat lagi")
)

func Authenticate() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return CheckAuthenticate(c)
	}
}

func CheckAuthenticate(c *fiber.Ctx) error {
	authorization := c.Request().Header.Peek("Authorization")
	if authorization != nil {
		splitToken := strings.Split(string(authorization), "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("server.secret_key")), nil
		})
		if err != nil {
			return responseErr.Response(c, responseErr.New(fiber.StatusUnauthorized, responseErr.WithMessage(ErrUnauthorized.Error())))
		}
		c.Locals("user", token)
		return c.Next()
	}
	c.Locals("user", nil)
	return c.Next()
}

func User(c context.Context) *JWTClaims {
	if c.Value("user") == nil {
		return nil
	}
	token := c.Value("user").(*jwt.Token)
	if !token.Valid {
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return &JWTClaims{
		ID:   uint64(claims["id"].(float64)),
	}
}

type JWTClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}
