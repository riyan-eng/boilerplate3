package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/app"
)

func TaskRoute(a *fiber.App, handler *app.ServiceServer) {
	route := a.Group("/task")
	route.Get("/", handler.ListTask)
	route.Post("/", handler.CreateTask)
	route.Get("/:id", handler.DetailTask)
	route.Patch("/:id", handler.UpdateTask)
	route.Delete("/:id", handler.DeleteTask)
}
