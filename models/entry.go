package models

import (
	"ginrestapi/database"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Bookname string `gorm:"type:text" json:"bookname"`
	UserID   uint
}

func (book *Book) Save() (*Book, error) {
	err := database.Database.Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}
