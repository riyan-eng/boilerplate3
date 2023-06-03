package config

import "github.com/gofiber/fiber/v2/middleware/cors"

func NewCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "http://127.0.0.1:3000, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}
}
