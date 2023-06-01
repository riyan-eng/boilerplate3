package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/route"
	"github.com/riyan-eng/boilerplate3/internal/service"
	// "github.com/riyan-eng/boilerplate3/internal/app"
)

func main() {
	// database
	db, err := repository.NewDB()
	if err != nil {
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
	route.NewRoute(fiberApp, taskService)
	// log.Println("he")
	// fiberApp.Get("/task", func(c *fiber.Ctx) error {
	// 	return c.SendString("halo")
	// })
	fiberApp.Listen(":3000")
}
