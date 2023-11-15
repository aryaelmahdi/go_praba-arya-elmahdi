package query

import "tugas/praktikum/model"

func (usersQuery *UserQueryImpl) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	result := usersQuery.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (usersQuery *UserQueryImpl) CreateUser(user model.User) (*model.User, error) {
	result := usersQuery.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
