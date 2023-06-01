package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) DeleteTask(c *fiber.Ctx) error {
	return c.SendString("delete task")
}
