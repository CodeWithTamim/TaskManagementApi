package repository

import (
	"github.com/CodeWithTamim/TaskManagementApi/config"
	"github.com/CodeWithTamim/TaskManagementApi/models"
)

func RegisterUser(email, password string) error {
	return nil
}

func GetUserProfile(email string) error {
	return nil
}

func UserExists(email string) bool {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return false
	}
	return true
}
