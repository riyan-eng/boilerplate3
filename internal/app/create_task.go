package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
	"github.com/riyan-eng/boilerplate3/pkg/validation"
)

// @Summary     Create akun
// @Tags       	Akun
// @Accept		json
// @Produce		json
// @Param       body	body  srv.TaskCreateReq	true  "body"
// @Router		/task [post]
func (s *ServiceServer) CreateTask(c *fiber.Ctx) error {
	body := new(srv.TaskCreateReq)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	validation.ValidateCreateTask(*body)

	data := s.taskService.CreateTask(dto.TaskCreateReq{
		ID:     uuid.NewString(),
		Name:   body.Name,
		Detail: body.Detail,
	})
	return c.JSON(util.Response{
		Data:    data.Data,
		Message: util.MESSAGE_OK_CREATE,
	})
}
