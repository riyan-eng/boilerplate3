package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/riyan-eng/boilerplate3/docs"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/route"
	"github.com/riyan-eng/boilerplate3/internal/service"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	// database
	db, err := repository.NewDB()
	if err != nil {
		log.Printf("cannot connect database: %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping database: %v", err)
	}

	// register service
	dao := repository.NewDAO(db)
	taskService := service.NewTaskService(dao)

	//
	fiberApp := fiber.New()
	fiberApp.Get("/docs/*", swagger.New(swagger.Config{
		Title: "BP3",
	}))
	route.NewRoute(fiberApp, taskService)
	// log.Println("he")
	// fiberApp.Get("/task", func(c *fiber.Ctx) error {
	// 	return c.SendString("halo")
	// })
	fiberApp.Listen(":3000")
}
