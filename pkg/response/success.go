package response

import (
	"github.com/gofiber/fiber/v2"
)

type AppSuccessOption func(*AppSuccess)

type AppSuccess struct {
	Data *interface{} `json:"data,omitempty"`
	Meta *interface{} `json:"meta,omitempty"`
}

type Message struct {
	Message *string `json:"message,omitempty"`
}

func SuccessMeta(meta interface{}) AppSuccessOption {
	if meta != nil {
		return func(h *AppSuccess) {
			h.Meta = &meta
		}
	}
	return nil
}
func SuccessData(data interface{}) AppSuccessOption {
	if data != nil {
		return func(h *AppSuccess) {
			h.Data = &data
		}
	}
	return nil
}

// SuccessMessage
func SuccessMessage(message *string) interface{} {
	if message != nil {
		return Message{
			Message: message,
		}
	}
	return nil
}


func Success(c *fiber.Ctx, status int, option ...AppSuccessOption) error {
	appSuccess := new(AppSuccess)
	// Loop through each option
	for _, opt := range option {
		// Call the option giving the instantiated
		opt(appSuccess)
	}
	if status == 200 {
		return c.Status(fiber.StatusOK).JSON(*appSuccess)
	} else if status > 200 && status < 300 {
		return c.Status(status).JSON(*appSuccess)
	}
	return nil
}
