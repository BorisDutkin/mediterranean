package main

import (
	"log"

	"github.com/BorisDutkin/mediterranean/database"
	"github.com/BorisDutkin/mediterranean/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()	

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}					

	routes.Setup(app)
	database.Connect()
	
	app.Listen(":3000")
}