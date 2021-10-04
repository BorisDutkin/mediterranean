package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/BorisDutkin/mediterranean/database"
	"github.com/BorisDutkin/mediterranean/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Register - new user registry
func Register(c *fiber.Ctx) error {		
	payload := struct {
		models.User
		PasswordConfirm string `json:"password_confirm"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}	

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password doesn't match",
		})
	}	

	user := models.User {
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: payload.Password,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	token, err := GenerateJwtClaims(user); if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})	
}

// Login - user login
func Login(c *fiber.Ctx) error {
	payload := struct {
		Email string
		Password string
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}	
	
	var user models.User

	database.DB.Where("email = ?", payload.Email).First(&user)

	if user.ID == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := user.ValidatePassword(payload.Password); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	token, err := GenerateJwtClaims(user); if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})	
}

// GenerateJwtClaims - generates jwt token with claims
func GenerateJwtClaims(u models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.FirstName + " " + u.LastName	
	claims["expiration"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", nil
	}

	return t, nil
}

func Hello(c *fiber.Ctx) error {	
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"]
	
	return c.JSON(fiber.Map{"message": fmt.Sprintf("You're in, %s!", name)})	
}