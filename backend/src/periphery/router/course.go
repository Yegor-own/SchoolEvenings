package router

import (
	"backend/src/domain/usecase"
	"backend/src/periphery/controller"
	"backend/src/periphery/middleware"
	"github.com/gofiber/fiber/v2"
)

func CourseRouter(app fiber.Router, service usecase.CourseService) {
	app.Get("/all", controller.GetAllCourses(service))
	app.Get("/byId", controller.GetCourse(service))
	app.Post("/create", middleware.ProtectJWT())
	app.Patch("/update", middleware.ProtectJWT())
	app.Delete("/delete", middleware.ProtectJWT())
	//app.Post("/login", controller.Login(service))
	//app.Post("/register", controller.RegisterUser(service))
	//app.Get("/getById", middleware.ProtectJWT(), controller.GetUser(service))
	//app.Patch("/update", middleware.ProtectJWT(), controller.UpdateUser(service))
	//app.Delete("/delete", middleware.ProtectJWT(), controller.DeleteUser(service))
}
