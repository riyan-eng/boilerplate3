package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/app"
)

func AuthRoute(a *fiber.App, handler *app.ServiceServer) {
	route := a.Group("/auth")
	route.Post("/register", handler.Register)
	route.Post("/login", handler.Login)
	route.Post("/refresh_token", handler.RefreshToken)
}
