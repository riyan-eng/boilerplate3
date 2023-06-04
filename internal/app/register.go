package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) Register(c *fiber.Ctx) error {
	s.authService.Register()
	return c.SendString("register")
}
