package controller

import (
	"backend/internal/gateway"
	"backend/internal/presenter"
	"backend/internal/rule"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(service rule.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params gateway.CreateUser
		err := ctx.BodyParser(&params)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.CreateUser(params.Name, params.Email)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(res))
	}
}
