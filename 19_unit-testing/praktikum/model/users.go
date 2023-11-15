package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blog     []Blog `gorm:"foreignKey:UserID;references:id"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInterface interface {
	Register(newUser Users) *Users
	GetAll() []Users
	GetByID(id int) *Users
	Delete(id int)
	Update(updatedData Users) bool
	Login(email string, password string) *Users
}

type UserModel struct {
	db *gorm.DB
}

func (u *UserModel) Init(db *gorm.DB) {
	u.db = db
}

func (u *UserModel) Register(newUser Users) *Users {
	if err := u.db.Create(&newUser).Error; err != nil {
		logrus.Error("Insert data failed", err.Error())
		return nil
	}
	return &newUser
}

func (u *UserModel) GetAll() []Users {
	listUser := []Users{}
	if err := u.db.Find(&listUser).Error; err != nil {
		logrus.Error("Model : failed to get data, ", err.Error())
		return nil
	}
	return listUser
}

func (u *UserModel) GetByID(id int) *Users {
	user := Users{}
	if err := u.db.Find(&user, "id = ?", id).Error; err != nil {
		logrus.Error("Model :failed to get user data, ", err.Error())
	}

	if uint(id) > user.ID || id < 0 {
		return nil
	}
	return &user
}

func (u *UserModel) Delete(id int) {
	var deletdData = Users{}
	deletdData.ID = uint(id)
	if err := u.db.Delete(&deletdData).Error; err != nil {
		logrus.Error("Model : Delete error, ", err.Error())
	}
}

func (u *UserModel) Update(updatedData Users) bool {
	query := u.db.Table("users").Where("id = ?", updatedData.ID).Updates(&updatedData)
	if err := query.Error; err != nil {
		logrus.Error("Model : update error, ", err.Error())
		return false
	}

	if dataCount := query.RowsAffected; dataCount < 1 {
		logrus.Error("Model : update error, not rows affected")
		return false
	}

	return true
}

func (u *UserModel) Login(email string, password string) *Users {
	var userData = Users{}
	if err := u.db.Where("email = ? AND password = ?", email, password).First(&userData).Error; err != nil {
		logrus.Error("Model : Login failed,", err.Error())
		return nil
	}

	if userData.ID == 0 {
		logrus.Error("Model : User not found")
		return nil
	}

	return &userData
}
