package router

import (
	"backend/src/domain/usecase"
	"backend/src/periphery/controller"
	"backend/src/periphery/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service usecase.UserService) {
	app.Post("/login", controller.Login(service))
	app.Post("/register", controller.RegisterUser(service))
	app.Get("/getById", middleware.ProtectJWT(), controller.GetUser(service))
	app.Patch("/update", middleware.ProtectJWT(), controller.UpdateUser(service))
	app.Delete("/delete", middleware.ProtectJWT(), controller.DeleteUser(service))
}
