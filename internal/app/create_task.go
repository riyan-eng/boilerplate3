package app

import "github.com/gofiber/fiber/v2"

// GetStringByInt example
//
//	@Summary		Add a new pet to the store
//	@Description	get string by ID
//	@ID				get-string-by-int
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	string			"ok"
//	@Router			/testapi/get-string-by-int/{some_id} [get]
func (s *ServiceServer) CreateTask(c *fiber.Ctx) error {
	return c.SendString("create task")
}
