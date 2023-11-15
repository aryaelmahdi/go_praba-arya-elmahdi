package routes

import (
	"tugas/praktikum/features/users/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UsersRoute(e echo.Echo, handler handler.UsersHandlerInterface, secret string) {
	e.POST("/users", handler.CreateUser(), echojwt.JWT([]byte(secret)))
	e.GET("/users", handler.GetAllUsers(), echojwt.JWT([]byte(secret)))
}
