package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

// @Summary      Delete akun
// @Tags       	 Akun
// @Produce      json
// @Param        id		path	string	true	"id"
// @Router       /task/{id} [delete]
func (s *ServiceServer) DeleteTask(c *fiber.Ctx) error {
	// cek existing data with id
	s.taskService.DetailTask(dto.TaskDetailReq{
		ID: c.Params("id"),
	})
	// service delete
	s.taskService.DeleteTask(dto.TaskDeleteReq{
		ID: c.Params("id"),
	})
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_DELETE,
	})
}
