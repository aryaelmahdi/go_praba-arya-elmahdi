package service

import (
	"tugas/praktikum/features/users/query"
	"tugas/praktikum/model"
)

type UsersServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) (*model.User, error)
}

type UsersServiceImpl struct {
	userQuery query.UsersQuery
}

func NewUsersServiceInterface(query query.UsersQuery) *UsersServiceImpl {
	return &UsersServiceImpl{userQuery: query}
}
