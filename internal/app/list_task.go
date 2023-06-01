package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) ListTask(c *fiber.Ctx) error {
	return c.SendString("list task")
}
