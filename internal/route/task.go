package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/app"
	"github.com/riyan-eng/boilerplate3/internal/service"
)

func TaskRoute(a *fiber.App, taskService service.TaskService) {
	handler := app.NewService(taskService)
	route := a.Group("/task")
	route.Get("/", handler.ListTask)
	route.Post("/", handler.CreateTask)
	route.Get("/:id", handler.DetailTask)
	route.Patch("/:id", handler.UpdateTask)
	route.Delete("/:id", handler.DeleteTask)
}
