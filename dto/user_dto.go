package dto

import (
	"time"

	"github.com/CodeWithTamim/TaskManagementApi/models"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=5,max=20"`
	Email    string `json:"email" validate:"required,email,min=5,max=20"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=5,max=20"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type RegisterResponse struct {
	Id        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginResponse struct {
	Id        uint
	Name      string
	Email     string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []models.Task
}
