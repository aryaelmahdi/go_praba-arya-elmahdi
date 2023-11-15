package service

import (
	"errors"
	"tugas/praktikum/model"
)

func (service *UsersServiceImpl) GetAllUsers() ([]model.User, error) {
	result, err := service.userQuery.GetAllUsers()
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return result, nil
}

func (service *UsersServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	result, err := service.userQuery.CreateUser(user)
	if err != nil {
		return nil, errors.New("wrong input")
	}
	return result, nil
}
