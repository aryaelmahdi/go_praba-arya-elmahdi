package main

import (
	"tugas/praktikum/config"
	"tugas/praktikum/features/users/handler"
	"tugas/praktikum/features/users/query"
	"tugas/praktikum/features/users/service"
	"tugas/praktikum/routes"
	"tugas/praktikum/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config := config.Init()
	db, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	userQuery := query.NewUsersModel(db)
	userService := service.NewUsersServiceInterface(userQuery)
	userHandler := handler.NewUsersHandler(userService)

	routes.UsersRoute(e, userHandler, config.Secret)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	e.Logger.Fatal(e.Start(":8080").Error())
}
