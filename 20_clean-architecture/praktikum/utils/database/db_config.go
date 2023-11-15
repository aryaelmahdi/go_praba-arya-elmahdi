package database

import (
	"errors"
	"tugas/praktikum/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("cannot connect to database, " + err.Error())
	}
	migrateDB(db)
	return db, nil
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
