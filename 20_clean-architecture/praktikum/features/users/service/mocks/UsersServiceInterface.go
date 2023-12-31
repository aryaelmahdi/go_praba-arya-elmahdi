// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "tugas/praktikum/model"

	mock "github.com/stretchr/testify/mock"
)

// UsersServiceInterface is an autogenerated mock type for the UsersServiceInterface type
type UsersServiceInterface struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UsersServiceInterface) CreateUser(user model.User) (*model.User, error) {
	ret := _m.Called(user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, error)); ok {
		return rf(&user)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(&user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(&user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields:
func (_m *UsersServiceInterface) GetAllUsers() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsersServiceInterface creates a new instance of UsersServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersServiceInterface {
	mock := &UsersServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
