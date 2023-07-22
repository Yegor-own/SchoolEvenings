package middleware

import (
	"backend/src/intermediary/output"
	"errors"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func ProtectJWT() fiber.Handler {

	// Get env salt
	err := godotenv.Load(".env")
	if err != nil {
		return func(ctx *fiber.Ctx) error {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: errors.Join(errors.New("Sorry smth went wrong"), err).Error(),
			})
		}
	}

	salt := []byte(os.Getenv("JWT"))
	if os.Getenv("JWT") == "" {
		return func(ctx *fiber.Ctx) error {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: errors.New("Sorry smth went wrong").Error(),
			})
		}
	}

	//TODO check if user is not him

	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return ctx.Status(fiber.StatusBadRequest).JSON(output.Output{
					Data:  nil,
					Error: err.Error(),
				})
			}
			return ctx.Status(fiber.StatusUnauthorized).JSON(output.Output{
				Data:  nil,
				Error: err.Error(),
			})
		},
		SigningKey: salt,
	})
}
