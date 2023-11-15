package routes

import (
	"praktikum/praktikum/config"
	"praktikum/praktikum/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController, config config.Config) {
	e.POST("/users", uc.Register())
	e.GET("/users", uc.GetAll())
	e.GET("/users/:id", uc.GetByID())
	e.PUT("/users/:id", uc.Update())
	e.DELETE("/users/:id", uc.Delete())
	e.POST("/users/login", uc.Login())
}

func BookRoutes(e *echo.Echo, bc controller.BookController, config config.Config) {
	e.POST("/books", bc.Create())
	e.GET("/books", bc.GetAllBooks())
	e.GET("/books/:id", bc.GetBookByID())
	e.PUT("/books/:id", bc.Update())
	e.DELETE("/books/:id", bc.DeleteBook())
}

func BlogRoutes(e *echo.Echo, blc controller.BlogController) {
	e.POST("/blogs", blc.CreateBlog())
	e.GET("/blogs", blc.GetAllBlogs())
	e.GET("/blogs/:id", blc.GetBlogById())
	e.PUT("/blogs/:id", blc.UpdateBlog())
	e.DELETE("/blogs/:id", blc.DeleteBlog())
}
