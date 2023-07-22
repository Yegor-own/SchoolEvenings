package controller

import (
	"backend/src/domain/usecase"
	"backend/src/intermediary/output"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAll(service usecase.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		courses, err := service.GetAll()
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  courses,
			Error: "",
		})
	}
}
