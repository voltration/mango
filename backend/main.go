package main

import (
	"mango/db"
	"mango/handlers"
	authHandler "mango/handlers/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	db.InitDatabase()
	app := fiber.New()
	app.Static("/", "./public")

	// for testing
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://127.0.0.1:3000, http://localhost:3000",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", handlers.RootHandler)

	// Auth Routes
	auth := app.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	app.Listen(":3000")
}
