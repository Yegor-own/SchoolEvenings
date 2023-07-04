package router

import (
	"backend/src/domain/usecase"
	"backend/src/periphery/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service usecase.UserService) {
	//app.Post("/login")
	app.Post("/register", controller.RegisterUser(service))
	app.Get("/getById", controller.GetUser(service))
	app.Patch("/update", controller.UpdateUser(service))
	app.Delete("/delete", controller.DeleteUser(service))
}
