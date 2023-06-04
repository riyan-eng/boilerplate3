package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
	"github.com/riyan-eng/boilerplate3/pkg/validation"
)

// @Summary     Update akun
// @Tags       	Akun
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Param       body	body	srv.TaskUpdateReq	true	"body"
// @Router      /task/{id} [patch]
func (s *ServiceServer) UpdateTask(c *fiber.Ctx) error {
	// cek existing data with id
	s.taskService.DetailTask(dto.TaskDetailReq{
		ID: c.Params("id"),
	})

	// parse & validate
	body := new(srv.TaskUpdateReq)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	validation.ValidateUpdateTask(*body)

	s.taskService.UpdateTask(dto.TaskUpdateReq{
		ID:     c.Params("id"),
		Name:   body.Name,
		Detail: body.Detail,
	})
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_UPDATE,
	})
}
