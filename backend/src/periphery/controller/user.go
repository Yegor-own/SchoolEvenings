package controller

import (
	"backend/src/domain/entity"
	"backend/src/domain/usecase"
	"backend/src/intermediary/input"
	"backend/src/intermediary/output"
	"backend/src/periphery/hash"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func RegisterUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.UserCreate
		err := ctx.BodyParser(&in)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Hash password
		password, err := hash.HashPassword(in.Password)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Insert User
		user, err := service.Register(in.Name, in.Surname, in.Patronymic, in.Email, in.Phone, password)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  *user,
			Error: nil,
		})
	}
}

func GetUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.UserById
		err := ctx.BodyParser(&in)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Fetch User by ID
		user, err := service.Get(in.ID)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  *user,
			Error: nil,
		})
	}
}

func UpdateUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in entity.User
		err := ctx.BodyParser(&in)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Update User
		user, err := service.ChangeData(in)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  *user,
			Error: nil,
		})
	}
}

func DeleteUser(service usecase.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Process input
		var in input.UserDelete
		err := ctx.BodyParser(&in)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Delete User
		err = service.Delete(in.ID)
		if err != nil {
			log.Println(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(output.Output{
				Data:  nil,
				Error: err,
			})
		}

		// Return User as a result
		ctx.Status(http.StatusOK)
		return ctx.JSON(output.Output{
			Data:  nil,
			Error: nil,
		})
	}
}
