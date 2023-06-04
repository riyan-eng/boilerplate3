package app

import "github.com/gofiber/fiber/v2"

func (s *ServiceServer) RefreshToken(c *fiber.Ctx) error {
	s.authService.RefreshToken()
	return c.SendString("refresh")
}
