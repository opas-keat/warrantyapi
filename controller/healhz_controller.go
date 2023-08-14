package controller

import (
	"warrantyapi/model"

	"github.com/gofiber/fiber/v2"
)

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Hello i'm ok!",
		Data:    nil,
	})
}

// Prepare an endpoint for 'Not Found'.
func NotFound(c *fiber.Ctx) error {
	return fiber.ErrServiceUnavailable
	// errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
	// return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
	// 	Code:    "999",
	// 	Message: errorMessage,
	// 	Data:    "",
	// })
}
