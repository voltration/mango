package main

import (
	"mango/db"
	authHandler "mango/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.InitDatabase()

	app := fiber.New()

	// Auth Routes
	auth := app.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	app.Listen(":3000")
}
