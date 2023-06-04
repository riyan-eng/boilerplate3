package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) Login(c *fiber.Ctx) error {
	s.authService.Login()
	return c.SendString("login")
}
