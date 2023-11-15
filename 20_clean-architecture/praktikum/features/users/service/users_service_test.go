package service

import (
	"errors"
	"testing"
	"tugas/praktikum/features/users/service/mocks"
	"tugas/praktikum/model"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	data := mocks.NewUsersQuery(t)
	var service = NewUsersServiceInterface(data)

	var newData = model.User{
		Password: "gugu",
		Email:    "arya@gmail.com",
	}

	invalidData := model.User{
		Password: "",
		Email:    "",
	}

	t.Run("success insert", func(t *testing.T) {
		data.On("CreateUser", newData).Return(&newData, nil).Once()
		result, err := service.CreateUser(newData)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Failed insert", func(t *testing.T) {
		data.On("CreateUser", invalidData).Return(nil, errors.New("insert process failed")).Once()
		result, err := service.CreateUser(invalidData)
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	data.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	data := mocks.NewUsersQuery(t)
	var service = NewUsersServiceInterface(data)

	users := make([]model.User, 0)

	t.Run("success insert", func(t *testing.T) {
		data.On("GetAllUsers").Return(users, nil).Once()
		result, err := service.GetAllUsers()
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Failed insert", func(t *testing.T) {
		data.On("GetAllUsers").Return(nil, errors.New("no data")).Once()
		result, err := service.GetAllUsers()
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	data.AssertExpectations(t)
}
