package query

import (
	"tugas/praktikum/model"

	"gorm.io/gorm"
)

type UsersQuery interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) (*model.User, error)
}

type UserQueryImpl struct {
	DB *gorm.DB
}

func NewUsersModel(db *gorm.DB) *UserQueryImpl {
	return &UserQueryImpl{DB: db}
}
