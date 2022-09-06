package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgustavopalmieri/go-admin-api/go/database"
	"github.com/lgustavopalmieri/go-admin-api/go/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)	

	app.Listen(":3000")
}
