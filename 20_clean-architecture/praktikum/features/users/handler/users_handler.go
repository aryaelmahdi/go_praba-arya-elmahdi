package handler

import (
	"tugas/praktikum/model"

	"github.com/labstack/echo/v4"
)

func (handler *UsersHandlerImpl) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := handler.userService.GetAllUsers()
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func (handler *UsersHandlerImpl) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}
