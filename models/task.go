package models

import "gorm.io/gorm"

type Task struct {
	//default
	//
	gorm.Model
	Title       string
	Description string
	IsCompleted bool
	UserId      uint
	User        User `gorm:"foreignKey:UserId;references:Id"`
}
