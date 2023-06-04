package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/app"
	"github.com/riyan-eng/boilerplate3/internal/service"
)

func NewRoute(
	fiberApp *fiber.App,
	taskService service.TaskService,
	authService service.AuthService,
) {
	allHandler := app.NewService(taskService, authService)
	TaskRoute(fiberApp, allHandler)
}
