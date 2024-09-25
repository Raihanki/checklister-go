package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidationErrorResponse struct {
	Error       string
	FailedField string
	Tag         string
}

func JsonResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	if status < 100 || status > 500 {
		log.Fatalf("invalid status code")
	}

	if status == 500 {
		message = "Internal Server Error"
	}

	return c.Status(status).JSON(ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
