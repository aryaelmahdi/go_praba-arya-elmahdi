package main

import (
	"fmt"
	"praktikum/praktikum/config"
	"praktikum/praktikum/controller"
	"praktikum/praktikum/model"
	"praktikum/praktikum/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config := config.InitConfig()
	db := model.InitModel(*config)
	model.Migrate(db)

	userModel := model.UserModel{}
	userModel.Init(db)
	bookModel := model.BookModel{}
	bookModel.Init(db)
	blogModel := model.BlogModel{}
	blogModel.Init(db)

	userController := controller.UserController{}
	userController.Init(userModel, *config)
	bookController := controller.BookController{}
	bookController.Init(bookModel)
	blogController := controller.BlogController{}
	blogController.Init(blogModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.UserRoutes(e, userController, *config)
	routes.BookRoutes(e, bookController, *config)
	routes.BlogRoutes(e, blogController)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
