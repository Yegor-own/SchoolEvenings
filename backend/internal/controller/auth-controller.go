package controller

import (
	"backend/internal/entity"
	"backend/internal/gateway"
	"backend/internal/hutils"
	"backend/internal/presenter"
	"backend/internal/rule"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(service rule.UserService) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		input := new(gateway.LoginInput)
		userData := new(entity.User)

		err := ctx.BodyParser(input)
		if err != nil {
			return ctx.JSON(presenter.BadResponse(err))
		}

		userData, err = service.GetUserByEmail(input.Email)
		if err != nil {
			return ctx.JSON(presenter.BadResponse(err))
		}

		if userData.Name == "" {
			return ctx.JSON(presenter.BadResponse(errors.New("there is no user found")))
		}

		log.Println(userData)

		if !hutils.ComparePassword(input.Password, userData.Password) {
			return ctx.JSON(presenter.BadResponse(errors.New("wrong hutils")))
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = userData.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.JSON(presenter.SuccessRequest(map[string]string{"message": "Success login", "token": t}))
		//return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
	}

}
