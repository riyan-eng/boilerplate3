package util

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: MESSAGE_BAD_REQUEST,
			Data:    err.Error(),
		})
	}
	return c.Status(fiber.StatusBadGateway).JSON(Response{
		Data:    err.Error(),
		Message: MESSAGE_BAD_SYSTEM,
	})
}
