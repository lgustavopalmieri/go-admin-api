package main

// example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lgustavopalmieri/go-admin-api/go/database"
	"github.com/lgustavopalmieri/go-admin-api/go/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	// that's important to frontend get the cookies
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8100")
}
