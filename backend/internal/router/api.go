package router

import (
	"backend/internal/database"
	"backend/internal/repo"
	"backend/internal/rule"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Api() {
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connection success")

	courseRepo := repo.NewCourseRepo(db)
	courseService := rule.NewCourseService(courseRepo)

	userRepo := repo.NewUserRepo(db)
	userService := rule.NewUserService(userRepo)

	app := fiber.New()
	app.Use(cors.New())

	courseRouter(app, courseService)
	userRouter(app, userService)

	log.Fatal(app.Listen(":8080"))
}
