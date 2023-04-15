package router

import (
	"backend/internal/controller"
	"backend/internal/rule"

	"github.com/gofiber/fiber/v2"
)

func userRouter(app *fiber.App, service rule.UserService) {
	app.Post("user/create", controller.CreateUser(service))
}
