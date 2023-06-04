package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

// @Summary      Detail akun
// @Tags       	 Akun
// @Produce      json
// @Param        id		path	string	true	"id"
// @Router       /task/{id} [get]
func (s *ServiceServer) DetailTask(c *fiber.Ctx) error {
	service := s.taskService.DetailTask(dto.TaskDetailReq{
		ID: c.Params("id"),
	})
	return c.JSON(util.Response{
		Data:    service.Item,
		Message: util.MESSAGE_OK_READ,
	})
}
