package main

import (
	"log"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/riyan-eng/boilerplate3/internal/repository"
	"github.com/riyan-eng/boilerplate3/internal/route"
	"github.com/riyan-eng/boilerplate3/internal/service"
	"github.com/riyan-eng/boilerplate3/pkg/docs"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}
}

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

	// swagger
	docs.SwaggerInfo.Title = "Boilerplate 3"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// fiber
	fiberApp := fiber.New()
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	fiberApp.Get("/docs/*", swagger.New(swagger.Config{
		Title: "BP3",
	}))
	route.NewRoute(fiberApp, taskService)
	fiberApp.Listen(":3000")
}
