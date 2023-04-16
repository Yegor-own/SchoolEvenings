package controller

import (
	"backend/internal/gateway"
	"backend/internal/presenter"
	"backend/internal/rule"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateUser(service rule.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params gateway.CreateUser
		err := ctx.BodyParser(&params)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		res, err := service.CreateUser(params.Name, params.Surname, params.Email, params.Password)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenter.BadResponse)
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = res.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenter.SuccessRequest(map[interface{}]interface{}{
			"user": res,
			"jwt":  t,
		}))
	}
}

func UpdateUser(service rule.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var
		return nil
	}
}
