package main

import (
	"mango/db"
	authHandler "mango/handlers/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	db.InitDatabase()

	app := fiber.New()

	app.Static("/", "/public")

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./public/login.html")
	})

	// for testing
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Auth Routes
	auth := app.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	app.Listen(":3000")
}
