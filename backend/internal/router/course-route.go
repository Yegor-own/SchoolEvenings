package router

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/rule"

	"github.com/gofiber/fiber/v2"
)

func courseRouter(app *fiber.App, service rule.CourseService) {
	app.Get("/course/getAll", controller.GetCourses(service))
	app.Post("/course/create", middleware.Protected(), controller.CreateCourse(service))
	app.Get("/course/getById", controller.GetCourseById(service))
}
