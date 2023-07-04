package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRouter() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(), logger.New())

	return app
}
