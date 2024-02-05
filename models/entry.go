package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Bookname string `gorm:"type:text" json:"bookname"`
	UserID   uint
}
