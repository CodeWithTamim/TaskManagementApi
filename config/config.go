package config

import (
	"fmt"
	"os"

	"github.com/CodeWithTamim/TaskManagementApi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	var err error
	databaseUser := os.Getenv("DATABASE_USER")
	databasePass := os.Getenv("DATABASE_PASS")
	databaseName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePass, databaseName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize database connection")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Task{})
}
