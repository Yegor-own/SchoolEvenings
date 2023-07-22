package main

import (
	"backend/src/domain/usecase"
	"backend/src/intermediary/repository"
	"backend/src/intermediary/storage"
	"backend/src/periphery/driver/sqlite"
	"backend/src/periphery/router"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", "database/backend.db")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	driver, err := sqlite.NewSqlxDriver(db, "./scripts/users.sql", "./scripts/courses.sql", "./scripts/registry.sql")
	if err != nil {
		log.Fatalln(err.Error())
	}
	access := storage.NewDataAccess(driver)
	userRepo := repository.NewUserRepository(access)
	userService := usecase.NewUserService(userRepo)

	app := router.NewRouter()
	api := app.Group("/api")
	user := api.Group("/user")
	router.UserRouter(user, userService)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
