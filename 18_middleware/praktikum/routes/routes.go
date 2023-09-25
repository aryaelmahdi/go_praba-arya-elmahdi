package routes

import (
	"praktikum/praktikum/config"
	"praktikum/praktikum/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController, config config.Config) {
	e.POST("/users", uc.Register())
	e.GET("/users", uc.GetAll(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/users/:id", uc.GetByID(), echojwt.JWT([]byte(config.Secret)))
	e.PUT("/users/:id", uc.Update(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/users/:id", uc.Delete(), echojwt.JWT([]byte(config.Secret)))
	e.POST("/users/login", uc.Login())
}

func BookRoutes(e *echo.Echo, bc controller.BookController, config config.Config) {
	e.POST("/books", bc.Create(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/books", bc.GetAllBooks(), echojwt.JWT([]byte(config.Secret)))
	e.GET("/books/:id", bc.GetBookByID(), echojwt.JWT([]byte(config.Secret)))
	e.PUT("/books/:id", bc.Update(), echojwt.JWT([]byte(config.Secret)))
	e.DELETE("/books/:id", bc.DeleteBook(), echojwt.JWT([]byte(config.Secret)))
}

func BlogRoutes(e *echo.Echo, blc controller.BlogController) {
	e.POST("/blogs", blc.CreateBlog())
	e.GET("/blogs", blc.GetAllBlogs())
	e.GET("/blogs/:id", blc.GetBlogById())
	e.PUT("/blogs/:id", blc.UpdateBlog())
	e.DELETE("/blogs/:id", blc.DeleteBlog())
}
