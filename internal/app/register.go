package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
	"github.com/riyan-eng/boilerplate3/pkg/validation"
)

// @Summary     Register
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  srv.AuthRegister	true  "body"
// @Router		/auth/register [post]
func (s *ServiceServer) Register(c *fiber.Ctx) error {
	body := new(srv.AuthRegister)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	validation.ValidateAuthRegister(*body)

	fmt.Println(body)
	s.authService.Register()
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_CREATE,
	})
}
