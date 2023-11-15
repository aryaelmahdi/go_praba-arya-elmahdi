package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Blog struct {
	ID     uint   `json:"id" gorm:"type:int(20);primaryKey;autoIncrement"`
	UserID int    `json:"uid"`
	Judul  string `json:"judul" gorm:"type:varchar(50)"`
	Konten string `json:"konten" gorm:"type:varchar(50)"`
}

type BlogInterface interface {
	CreateBlog(blog Blog) *Blog
	GetAllBlogs() []Blog
	GetBlogByID(id int) *Blog
	DeleteBlog(id int)
	UpdateBlog(updatedData Blog) bool
}

type BlogModel struct {
	db *gorm.DB
}

func (b *BlogModel) Init(db *gorm.DB) {
	b.db = db
}

func (b *BlogModel) CreateBlog(blog Blog) *Blog {
	if err := b.db.Create(&blog).Error; err != nil {
		logrus.Error("Insert data failed", err.Error())
		return nil
	}
	return &blog
}

func (b *BlogModel) GetAllBlogs() []Blog {
	blogList := []Blog{}
	if err := b.db.Find(&blogList).Error; err != nil {
		logrus.Error("Model : failed to get data, ", err.Error())
		return nil
	}
	return blogList
}

func (b *BlogModel) GetBlogByID(id int) *Blog {
	blog := Blog{}
	if err := b.db.Find(&blog, "id = ?", id).Error; err != nil {
		logrus.Error("Model :failed to get user data, ", err.Error())
	}

	if uint(id) > blog.ID || id < 0 {
		return nil
	}
	return &blog
}

func (b *BlogModel) DeleteBlog(id int) {
	var deletdData = Blog{}
	deletdData.ID = uint(id)
	if err := b.db.Delete(&deletdData).Error; err != nil {
		logrus.Error("Model : Delete error, ", err.Error())
	}
}

func (b *BlogModel) UpdateBlog(updatedData Blog) bool {
	query := b.db.Table("blogs").Where("id = ?", updatedData.ID).Updates(&updatedData)
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
