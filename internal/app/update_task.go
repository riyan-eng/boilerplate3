package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) UpdateTask(c *fiber.Ctx) error {
	return c.SendString("update task")
}
