package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/app"
	"github.com/riyan-eng/boilerplate3/internal/service"
)

func TaskRoute(a *fiber.App, taskService service.TaskService) {
	handler := app.NewService(taskService)
	route := a.Group("/task")
	route.Get("/", handler.CreateTask)
	route.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("ini detail")
	})
	route.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("ini create")
	})
	route.Put("/", func(c *fiber.Ctx) error {
		return c.SendString("ini update")
	})
	route.Delete("/", func(c *fiber.Ctx) error {
		return c.SendString("ini delete")
	})
}
