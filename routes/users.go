package routes

import (
	"github.com/BorisDutkin/mediterranean/controllers"
	"github.com/BorisDutkin/mediterranean/middlewares"
	"github.com/gofiber/fiber/v2"
)

// SetUsersRoutes - users routing setup
func SetUsersRoutes(r fiber.Router) {
	users := r.Group("/users");	
	users.Post("/register", controllers.Register)	
	users.Post("/login", controllers.Login)	
	users.Get("/test", middlewares.Protected(), controllers.Hello)	
}