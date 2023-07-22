package controller

import (
	"backend/src/domain/entity"
	"backend/src/domain/usecase"
	"backend/src/intermediary/input"
	"backend/src/intermediary/output"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func getJWTClaims(ctx *fiber.Ctx) (jwt.MapClaims, error) {
	headers := ctx.GetReqHeaders()
	tokenString, ok := headers["Authorisation"]
	if !ok {
		err := errors.New("Sorry no token")
		log.Println(err.Error())
		return nil, err
	}

	tokenString = strings.Split(tokenString, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		err := godotenv.Load(".env")
		return nil, err
		salt := []byte(os.Getenv("JWT"))

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return salt, nil
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err := errors.New("Sorry bad claims")
		log.Println(err.Error())
		return nil, err
	}

	return claims, nil
}

func handleErrors(ctx *fiber.Ctx, err error) error {
	log.Println(err)
	ctx.Status(http.StatusBadRequest)
	return ctx.JSON(output.Output{
		Data:  nil,
		Error: err.Error(),
	})
}

func Login(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.Login
		err := ctx.BodyParser(&in)
		if err != nil {
			return handleErrors(ctx, err)
		}

		token, err := service.Login(in.Email, in.Password)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return JWT token as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  token,
			Error: "",
		})
	}
}

func RegisterUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.UserCreate
		err := ctx.BodyParser(&in)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Insert User
		user, err := service.Register(in.Name, in.Surname, in.Patronymic, in.Email, in.Phone, in.Password)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Login User
		token, err := service.Login(in.Email, in.Password)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data: map[string]any{
				"user":  *user,
				"token": token,
			},
			Error: "",
		})
	}
}

func GetUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get claims from Header
		claims, err := getJWTClaims(ctx)
		email, ok := claims["email"].(string)
		if !ok {
			err = errors.New("No email claims")
			return handleErrors(ctx, err)
		}

		// Fetch User by Email
		user, err := service.GetByEmail(email)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  *user,
			Error: "",
		})
	}
}

func UpdateUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.UserUpdate
		err := ctx.BodyParser(&in)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Get claims
		claims, err := getJWTClaims(ctx)
		email, ok := claims["email"].(string)
		if !ok {
			err = errors.New("No email claims")
			return handleErrors(ctx, err)
		}

		// Fetch User by Email
		user, err := service.GetByEmail(email)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Update User
		user, err = service.ChangeData(entity.ChangeUserData{
			ID:         user.ID,
			Name:       in.Name,
			Surname:    in.Surname,
			Patronymic: in.Patronymic,
			Phone:      in.Phone,
		})
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  *user,
			Error: "",
		})
	}
}

func DeleteUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get claims
		claims, err := getJWTClaims(ctx)
		email, ok := claims["email"].(string)
		if !ok {
			err = errors.New("No email claims")
			return handleErrors(ctx, err)
		}

		// Fetch User by Email
		user, err := service.GetByEmail(email)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Delete User
		err = service.Delete(user.ID)
		if err != nil {
			return handleErrors(ctx, err)
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  nil,
			Error: "",
		})
	}
}
