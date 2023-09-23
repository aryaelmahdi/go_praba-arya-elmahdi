package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Books struct {
	ID       uint   `json:"id" gorm:"type:int(20);primaryKey;autoIncrement"`
	Nama     string `json:"nama" gorm:"type:varchar(50)"`
	Penulis  string `json:"penulis" gorm:"type:varchar(50)"`
	Penerbit string `json:"penerbit" gorm:"type:varchar(50)"`
}

type BookModel struct {
	db *gorm.DB
}

func (b *BookModel) Init(db *gorm.DB) {
	b.db = db
}

func (b *BookModel) Insert(book Books) *Books {
	if err := b.db.Create(&book).Error; err != nil {
		logrus.Error("Insert data failed", err.Error())
		return nil
	}
	return &book
}

func (b *BookModel) GetAllBooks() []Books {
	bookList := []Books{}
	if err := b.db.Find(&bookList).Error; err != nil {
		logrus.Error("Model : failed to get data, ", err.Error())
		return nil
	}
	return bookList
}

func (b *BookModel) GetBookByID(id int) *Books {
	book := Books{}
	if err := b.db.Find(&book, "id = ?", id).Error; err != nil {
		logrus.Error("Model :failed to get user data, ", err.Error())
	}

	if uint(id) > book.ID || id < 0 {
		return nil
	}
	return &book
}

func (b *BookModel) DeleteBook(id int) {
	var deletdData = Books{}
	deletdData.ID = uint(id)
	if err := b.db.Delete(&deletdData).Error; err != nil {
		logrus.Error("Model : Delete error, ", err.Error())
	}
}

func (b *BookModel) UpdateBook(updatedData Books) bool {
	query := b.db.Table("books").Where("id = ?", updatedData.ID).Updates(&updatedData)
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
