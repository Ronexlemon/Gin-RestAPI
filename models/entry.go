package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName string `gorm:"type:text" json: "bookname"`
	UserID   uint
}
