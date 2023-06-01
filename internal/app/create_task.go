package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) CreateTask(c *fiber.Ctx) error {
	return c.SendString("create task")
}
