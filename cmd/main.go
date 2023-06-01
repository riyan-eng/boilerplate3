package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// database
	// db, err := repository.NewDB()
	// if err != nil {
	// 	return
	// }
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("cannot ping database: %v", err)
	// }

	// // register service
	// dao := repository.NewDAO(db)
	// service.NewTaskService(dao)

	//
	fiberApp := fiber.New()
	fiberApp.Listen(":3000")
}
