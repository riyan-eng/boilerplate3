package app

import "github.com/gofiber/fiber/v2"

// @Summary		Add a new pet to the store
// @Description	create string by ID
// @ID				create-string-by-int
// @Accept			json
// @Produce		json
// @Router			/task [post]
func (s *ServiceServer) CreateTask(c *fiber.Ctx) error {
	return c.SendString("create task")
}
