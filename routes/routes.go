package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {

	app := fiber.New()

	//define the routes
	//
	api := app.Group("/api")
	v1 := api.Group("/v1")

	{
		//User module
		v1.Post("/register")
		v1.Post("/login")
		v1.Get("/profile")
	}

	{
		//Task Module
		//=
		v1.Post("/tasks")
		v1.Get("/tasks")
		v1.Get("/tasks/:id")
		v1.Put("/tasks/:id")
		v1.Delete("/tasks/:id")
	}

	err := app.Listen(":3000")
	if err != nil {
		log.Printf("Error starting server, %v", err)
	}
}
