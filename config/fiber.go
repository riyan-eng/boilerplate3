package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/boilerplate3/pkg/util"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: util.ErrorHandler,
	}
}
