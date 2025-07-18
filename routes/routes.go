package routes

import (
	"log"

	"github.com/CodeWithTamim/TaskManagementApi/controllers"
	"github.com/CodeWithTamim/TaskManagementApi/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Run() {

	app := fiber.New()

	//define the routes
	//
	api := app.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")

	{
		//User module
		auth.Post("/register", controllers.RegisterHandler)
		auth.Post("/login", middlewares.AuthHandler, controllers.LoginHandler)
		auth.Get("/profile", middlewares.AuthHandler, middlewares.AuthHandler)
	}

	task := v1.Group("/task", middlewares.AuthHandler)
	{
		//Task Module
		//=
		task.Post("/tasks", controllers.AddTaskHandler)
		task.Get("/tasks", controllers.GetTasksHandler)
		task.Get("/tasks/:id", controllers.GetTasksHandler)
		task.Put("/tasks/:id", controllers.UpdateTaskHandler)
		task.Delete("/tasks/:id", controllers.DeleteTaskHandler)
	}

	err := app.Listen(":3000")

	if err != nil {
		log.Printf("Error starting server, %v", err)
	}
}
