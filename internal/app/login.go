package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
	"github.com/riyan-eng/boilerplate3/pkg/validation"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  srv.AuthLoginReq	true  "body"
// @Router		/auth/login [post]
func (s *ServiceServer) Login(c *fiber.Ctx) error {
	body := new(srv.AuthLoginReq)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	validation.ValidateAuthLogin(*body)
	service := s.authService.Login(dto.AuthLoginReq{
		Username: body.Username,
		Password: body.Password,
	})
	if !service.Match {
		return c.Status(fiber.StatusBadRequest).JSON(util.Response{
			Message: util.MESSAGE_FAILED_LOGIN,
		})
	}
	return c.JSON(util.Response{
		Data: fiber.Map{
			"user":          service.Data,
			"access_token":  service.AccessToken,
			"refresh_token": service.RefreshToken,
			"expired_at":    service.ExpiredAt,
		},
		Message: util.MESSAGE_OK_LOGIN,
	})
}
