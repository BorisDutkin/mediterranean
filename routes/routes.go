package routes

import "github.com/gofiber/fiber/v2"

// Setup - application routing setup
func Setup(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetUsersRoutes(v1)
}