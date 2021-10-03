package main

import (
	"github.com/BorisDutkin/mediterranean/database"
	"github.com/BorisDutkin/mediterranean/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()	

	routes.Setup(app)
	database.Connect()
	
	app.Listen(":3000")
}