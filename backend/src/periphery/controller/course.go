package controller

import (
	"backend/src/domain/usecase"
	"backend/src/intermediary/input"
	"backend/src/intermediary/output"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAllCourses(service usecase.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		courses, err := service.GetAll()
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return Course as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  courses,
			Error: "",
		})
	}
}

func GetCourse(service usecase.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.CourseById
		err := ctx.BodyParser(&in)
		if err != nil {
			return handleErrors(ctx, err)
		}

		courses, err := service.Get(in.ID)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return Course as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  courses,
			Error: "",
		})
	}
}

func UpdateCourse(service usecase.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

	}
}
