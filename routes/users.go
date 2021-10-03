package routes

import (
	"github.com/BorisDutkin/mediterranean/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetUsersRoutes - users routing setup
func SetUsersRoutes(r fiber.Router) {
	r.Get("/users/register", controllers.Register)	
}