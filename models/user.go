package models

import "time"

type User struct {
	//default
	//
	Id        uint `gorm:"primarykey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}
