package response

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrBadRequest = errors.New("your request is in a bad format")
	ErrInput      = errors.New("data is invalid")
	ErrDefault    = errors.New("server busy, please try again later")
	ErrUnprocessableEntity = errors.New("data is invalid")
)

type AppErrorOption func(*AppError)

// APPError is the default error struct containing detailed information about the error
type AppError struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Message is the error message that may be displayed to end users
	Message *string `json:"message,omitempty"`
	// Meta is the error detail detail data
	Meta *interface{} `json:"meta,omitempty"`
}

// New generates an application error
func New(status int, opts ...AppErrorOption) *AppError {
	err := new(AppError)
	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(err)
	}
	err.Status = status
	return err
}

// Error returns the error message.
func (e AppError) Error() string {
	return *e.Message
}

func WithMessage(message string) AppErrorOption {
	return func(h *AppError) {
		h.Message = &message
	}
}

func WithMeta(meta interface{}) AppErrorOption {
	return func(h *AppError) {
		h.Meta = &meta
	}
}

// Response writes an error response to client
func Response(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *AppError:
		return c.Status(e.Status).JSON(e)
	case validation.Errors:
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
}
