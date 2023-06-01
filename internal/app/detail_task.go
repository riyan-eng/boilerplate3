package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) DetailTask(c *fiber.Ctx) error {
	return c.SendString("detail task")
}
