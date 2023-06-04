package util

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// validation
	_, ok := err.(ValidationError)
	if ok {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Data:    err.Error(),
			Message: MESSAGE_BAD_REQUEST,
		})
	}
	// no data
	none := err == sql.ErrNoRows
	if none {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Data:    err.Error(),
			Message: MESSAGE_NOT_FOUND,
		})
	}
	return c.Status(fiber.StatusBadGateway).JSON(Response{
		Data:    err.Error(),
		Message: MESSAGE_BAD_SYSTEM,
	})
}
