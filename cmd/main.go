package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/route"
	"github.com/riyan-eng/boilerplate3/internal/service"
	"github.com/riyan-eng/boilerplate3/pkg/docs"
)

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

	// docs
	docs.SwaggerInfo.Title = "Boilerplate 3"
	// docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a sample swagger for Fiber"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//
	fiberApp := fiber.New()
	fiberApp.Get("/docs/*", swagger.New(swagger.Config{
		Title: "BP3",
	}))
	route.NewRoute(fiberApp, taskService)
	fiberApp.Listen(":3000")
}
