package main

import (
	"github.com/CodeWithTamim/TaskManagementApi/config"
	"github.com/CodeWithTamim/TaskManagementApi/routes"
	"github.com/CodeWithTamim/TaskManagementApi/utils"
)

func init() {
	utils.LoadEnv()
	config.InitDB()
	config.AutoMigrate()
}

func main() {

	token, _ := utils.CreateJWTClaim("tamimh.dev@gmail.com")
	println(token)
	routes.Run()
}
