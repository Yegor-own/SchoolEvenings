package controller

import (
	"backend/internal/gateway"
	"backend/internal/presenter"
	"backend/internal/rule"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetCourses(service rule.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var limit gateway.GetCourses
		err := ctx.BodyParser(&limit)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.FetchCourses(limit.Limit)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(res))
	}
}

func CreateCourse(service rule.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params gateway.CreateCourse
		err := ctx.BodyParser(&params)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.CreateCourse(params.Title, params.Description)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(res))
	}
}

func GetCourseById(service rule.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var id gateway.GetCourseById
		err := ctx.BodyParser(&id)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.GetCourseById(id.Id)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(res))
	}
}
