package app

import "github.com/gofiber/fiber/v2"

// @Summary      List akun
// @Tags       	 Akun
// @Accept       json
// @Produce      json
// @Param        order		query   string	false  "desc/asc default(desc)"
// @Param        search		query   string	false  "search"
// @Param        scope		query   string	false  "JU -> untuk jurnal umum"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per page"
// @Router       /task [get]
// @Security ApiKeyAuth
func (s *ServiceServer) ListTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "list task",
	})
}
