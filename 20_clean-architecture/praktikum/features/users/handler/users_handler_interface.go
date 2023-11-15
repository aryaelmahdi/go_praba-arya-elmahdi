package handler

import (
	"tugas/praktikum/features/users/service"

	"github.com/labstack/echo/v4"
)

type UsersHandlerInterface interface {
	CreateUser() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
}

type UsersHandlerImpl struct {
	userService service.UsersServiceInterface
}

func NewUsersHandler(service service.UsersServiceInterface) *UsersHandlerImpl {
	return &UsersHandlerImpl{userService: service}
}
