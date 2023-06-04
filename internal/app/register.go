package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
	"github.com/riyan-eng/boilerplate3/pkg/validation"
)

// @Summary     Register
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  srv.AuthRegisterReq	true  "body"
// @Router		/auth/register [post]
func (s *ServiceServer) Register(c *fiber.Ctx) error {
	body := new(srv.AuthRegisterReq)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	validation.ValidateAuthRegister(*body)

	s.authService.Register(dto.AuthRegisterReq{
		UserID:     uuid.NewString(),
		UserDataID: uuid.NewString(),
		Username:   body.Username,
		Email:      body.Email,
		Password:   body.Password,
		RoleCode:   body.RoleCode,
		Name:       body.Name,
		Gender:     body.Gender,
		Address:    body.Address,
		Phone:      body.Phone,
	})
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_CREATE,
	})
}
