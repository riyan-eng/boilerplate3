package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/dto"
	srv "github.com/riyan-eng/boilerplate3/pkg"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

// @Summary      List akun
// @Tags       	 Akun
// @Produce      json
// @Param        order		query   string	false  "desc/asc default(desc)"
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        limit		query   int		false  "limit"
// @Router       /task [get]
// @Security ApiKeyAuth
func (s *ServiceServer) ListTask(c *fiber.Ctx) error {
	queryParam := new(srv.TaskListReq).Init()
	err := c.QueryParser(&queryParam)
	util.PanicIfNeeded(err)

	service := s.taskService.ListTask(dto.TaskListReq{
		Search: queryParam.Search,
		Page:   queryParam.Page,
		Limit:  queryParam.Limit,
		Order:  queryParam.Order,
	})
	return c.JSON(util.Response{
		Data: service.Items,
		Meta: util.Meta{
			Page:  service.Page,
			Limit: service.Limit,
			Total: service.Total,
		},
		Message: util.MESSAGE_OK_READ,
	})
}
